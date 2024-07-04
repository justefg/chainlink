package client

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	pkgerrors "github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	commonassets "github.com/smartcontractkit/chainlink-common/pkg/assets"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"

	commonclient "github.com/smartcontractkit/chainlink/v2/common/client"
	commontypes "github.com/smartcontractkit/chainlink/v2/common/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/config"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
)

var (
	promEVMPoolRPCNodeDials = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "evm_pool_rpc_node_dials_total",
		Help: "The total number of dials for the given RPC node",
	}, []string{"evmChainID", "nodeName"})
	promEVMPoolRPCNodeDialsFailed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "evm_pool_rpc_node_dials_failed",
		Help: "The total number of failed dials for the given RPC node",
	}, []string{"evmChainID", "nodeName"})
	promEVMPoolRPCNodeDialsSuccess = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "evm_pool_rpc_node_dials_success",
		Help: "The total number of successful dials for the given RPC node",
	}, []string{"evmChainID", "nodeName"})

	promEVMPoolRPCNodeCalls = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "evm_pool_rpc_node_calls_total",
		Help: "The approximate total number of RPC calls for the given RPC node",
	}, []string{"evmChainID", "nodeName"})
	promEVMPoolRPCNodeCallsFailed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "evm_pool_rpc_node_calls_failed",
		Help: "The approximate total number of failed RPC calls for the given RPC node",
	}, []string{"evmChainID", "nodeName"})
	promEVMPoolRPCNodeCallsSuccess = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "evm_pool_rpc_node_calls_success",
		Help: "The approximate total number of successful RPC calls for the given RPC node",
	}, []string{"evmChainID", "nodeName"})
	promEVMPoolRPCCallTiming = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "evm_pool_rpc_node_rpc_call_time",
		Help: "The duration of an RPC call in nanoseconds",
		Buckets: []float64{
			float64(50 * time.Millisecond),
			float64(100 * time.Millisecond),
			float64(200 * time.Millisecond),
			float64(500 * time.Millisecond),
			float64(1 * time.Second),
			float64(2 * time.Second),
			float64(4 * time.Second),
			float64(8 * time.Second),
		},
	}, []string{"evmChainID", "nodeName", "rpcHost", "isSendOnly", "success", "rpcCallName"})
)

const rpcSubscriptionMethodNewHeads = "newHeads"

type rawclient struct {
	rpc  *rpc.Client
	geth *ethclient.Client
	uri  url.URL
}

type RpcClient struct {
	cfg     config.NodePool
	rpcLog  logger.SugaredLogger
	name    string
	id      int32
	chainID *big.Int
	tier    commonclient.NodeTier

	ws   rawclient
	http *rawclient

	stateMu sync.RWMutex // protects state* fields

	// Need to track subscriptions because closing the RPC does not (always?)
	// close the underlying subscription
	subs []ethereum.Subscription

	// Need to track the aliveLoop subscription, so we do not cancel it when checking lease on the MultiNode
	aliveLoopSub ethereum.Subscription

	// chStopInFlight can be closed to immediately cancel all in-flight requests on
	// this RpcClient. Closing and replacing should be serialized through
	// stateMu since it can happen on state transitions as well as RpcClient Close.
	chStopInFlight chan struct{}

	// intercepted values seen by callers of the rpcClient excluding health check calls. Need to ensure MultiNode provides repeatable read guarantee
	highestUserObservations commonclient.ChainInfo
	// most recent chain info observed during current lifecycle (reseted on DisconnectAll)
	latestChainInfo commonclient.ChainInfo
}

var _ commonclient.RPCClient[*big.Int, *evmtypes.Head] = &RpcClient{}

func NewRPCClient(
	cfg config.NodePool,
	lggr logger.Logger,
	wsuri url.URL,
	httpuri *url.URL,
	name string,
	id int32,
	chainID *big.Int,
	tier commonclient.NodeTier,
) *RpcClient {
	r := new(RpcClient)
	r.cfg = cfg
	r.name = name
	r.id = id
	r.chainID = chainID
	r.tier = tier
	r.ws.uri = wsuri
	if httpuri != nil {
		r.http = &rawclient{uri: *httpuri}
	}
	r.chStopInFlight = make(chan struct{})
	lggr = logger.Named(lggr, "Client")
	lggr = logger.With(lggr,
		"clientTier", tier.String(),
		"clientName", name,
		"client", r.String(),
		"evmChainID", chainID,
	)
	r.rpcLog = logger.Sugared(lggr).Named("RPC")

	return r
}

func (r *RpcClient) SubscribeToHeads(ctx context.Context) (<-chan *evmtypes.Head, commontypes.Subscription, error) {
	ctx, cancel, chStopInFlight, _, _ := r.acquireQueryCtx(ctx)
	defer cancel()

	newChainIDSubForwarder := func(chainID *big.Int, ch chan<- *evmtypes.Head) *subForwarder[*evmtypes.Head] {
		return newSubForwarder(ch, func(head *evmtypes.Head) *evmtypes.Head {
			head.EVMChainID = ubig.New(chainID)
			r.onNewHead(ctx, chStopInFlight, head)
			return head
		}, r.wrapRPCClientError)
	}

	ch := make(chan *evmtypes.Head)
	forwarder := newChainIDSubForwarder(r.chainID, ch)

	sub, err := r.subscribe(ctx, forwarder.srcCh, rpcSubscriptionMethodNewHeads)

	err = forwarder.start(sub, err)
	if err != nil {
		return nil, nil, err
	}

	err = r.registerSub(forwarder, r.chStopInFlight)
	if err != nil {
		return nil, nil, err
	}

	return ch, forwarder, err
}

func (r *RpcClient) SubscribeToFinalizedHeads(_ context.Context) (<-chan *evmtypes.Head, commontypes.Subscription, error) {
	interval := r.cfg.FinalizedBlockPollInterval()
	if interval == 0 {
		return nil, nil, errors.New("FinalizedBlockPollInterval is 0")
	}
	timeout := interval
	poller, channel := commonclient.NewPoller[*evmtypes.Head](interval, r.LatestFinalizedBlock, timeout, r.rpcLog)
	if err := poller.Start(); err != nil {
		return nil, nil, err
	}
	return channel, &poller, nil
}

func (r *RpcClient) Ping(ctx context.Context) error {
	version, err := r.ClientVersion(ctx)
	if err != nil {
		return fmt.Errorf("ping failed: %v", err)
	}
	r.rpcLog.Debugf("ping client version: %s", version)
	return err
}

func (r *RpcClient) UnsubscribeAllExcept(subs ...commontypes.Subscription) {
	r.stateMu.Lock()
	defer r.stateMu.Unlock()
	for _, sub := range r.subs {
		var keep bool
		for _, s := range subs {
			if sub == s {
				keep = true
				break
			}
		}
		if !keep {
			sub.Unsubscribe()
		}
	}
	r.latestChainInfo = commonclient.ChainInfo{}
}

// Not thread-safe, pure dial.
func (r *RpcClient) Dial(callerCtx context.Context) error {
	ctx, cancel := makeQueryCtx(callerCtx, r.getChStopInflight())
	defer cancel()

	promEVMPoolRPCNodeDials.WithLabelValues(r.chainID.String(), r.name).Inc()
	lggr := r.rpcLog.With("wsuri", r.ws.uri.Redacted())
	if r.http != nil {
		lggr = lggr.With("httpuri", r.http.uri.Redacted())
	}
	lggr.Debugw("RPC dial: evmclient.Client#dial")

	wsrpc, err := rpc.DialWebsocket(ctx, r.ws.uri.String(), "")
	if err != nil {
		promEVMPoolRPCNodeDialsFailed.WithLabelValues(r.chainID.String(), r.name).Inc()
		return r.wrapRPCClientError(pkgerrors.Wrapf(err, "error while dialing websocket: %v", r.ws.uri.Redacted()))
	}

	r.ws.rpc = wsrpc
	r.ws.geth = ethclient.NewClient(wsrpc)

	if r.http != nil {
		if err := r.DialHTTP(); err != nil {
			return err
		}
	}

	promEVMPoolRPCNodeDialsSuccess.WithLabelValues(r.chainID.String(), r.name).Inc()

	return nil
}

// Not thread-safe, pure dial.
// DialHTTP doesn't actually make any external HTTP calls
// It can only return error if the URL is malformed.
func (r *RpcClient) DialHTTP() error {
	promEVMPoolRPCNodeDials.WithLabelValues(r.chainID.String(), r.name).Inc()
	lggr := r.rpcLog.With("httpuri", r.ws.uri.Redacted())
	lggr.Debugw("RPC dial: evmclient.Client#dial")

	var httprpc *rpc.Client
	httprpc, err := rpc.DialHTTP(r.http.uri.String())
	if err != nil {
		promEVMPoolRPCNodeDialsFailed.WithLabelValues(r.chainID.String(), r.name).Inc()
		return r.wrapRPCClientError(pkgerrors.Wrapf(err, "error while dialing HTTP: %v", r.http.uri.Redacted()))
	}

	r.http.rpc = httprpc
	r.http.geth = ethclient.NewClient(httprpc)

	promEVMPoolRPCNodeDialsSuccess.WithLabelValues(r.chainID.String(), r.name).Inc()

	return nil
}

func (r *RpcClient) Close() {
	defer func() {
		if r.ws.rpc != nil {
			r.ws.rpc.Close()
		}
	}()
	r.cancelInflightRequests()
}

// cancelInflightRequests closes and replaces the chStopInFlight
func (r *RpcClient) cancelInflightRequests() {
	r.stateMu.Lock()
	defer r.stateMu.Unlock()
	close(r.chStopInFlight)
	r.chStopInFlight = make(chan struct{})
}

func (r *RpcClient) String() string {
	s := fmt.Sprintf("(%s)%s:%s", r.tier.String(), r.name, r.ws.uri.Redacted())
	if r.http != nil {
		s = s + fmt.Sprintf(":%s", r.http.uri.Redacted())
	}
	return s
}

func (r *RpcClient) logResult(
	lggr logger.Logger,
	err error,
	callDuration time.Duration,
	rpcDomain,
	callName string,
	results ...interface{},
) {
	lggr = logger.With(lggr, "duration", callDuration, "rpcDomain", rpcDomain, "callName", callName)
	promEVMPoolRPCNodeCalls.WithLabelValues(r.chainID.String(), r.name).Inc()
	if err == nil {
		promEVMPoolRPCNodeCallsSuccess.WithLabelValues(r.chainID.String(), r.name).Inc()
		logger.Sugared(lggr).Tracew(fmt.Sprintf("evmclient.Client#%s RPC call success", callName), results...)
	} else {
		promEVMPoolRPCNodeCallsFailed.WithLabelValues(r.chainID.String(), r.name).Inc()
		lggr.Debugw(
			fmt.Sprintf("evmclient.Client#%s RPC call failure", callName),
			append(results, "err", err)...,
		)
	}
	promEVMPoolRPCCallTiming.
		WithLabelValues(
			r.chainID.String(),             // chain id
			r.name,                         // RpcClient name
			rpcDomain,                      // rpc domain
			"false",                        // is send only
			strconv.FormatBool(err == nil), // is successful
			callName,                       // rpc call name
		).
		Observe(float64(callDuration))
}

func (r *RpcClient) getRPCDomain() string {
	if r.http != nil {
		return r.http.uri.Host
	}
	return r.ws.uri.Host
}

// registerSub adds the sub to the rpcClient list
func (r *RpcClient) registerSub(sub ethereum.Subscription, stopInFLightCh chan struct{}) error {
	r.stateMu.Lock()
	defer r.stateMu.Unlock()
	// ensure that the `sub` belongs to current life cycle of the `rpcClient` and it should not be killed due to
	// previous `DisconnectAll` call.
	select {
	case <-stopInFLightCh:
		sub.Unsubscribe()
		return fmt.Errorf("failed to register subscription - all in-flight requests were canceled")
	default:
	}
	// TODO: BCI-3358 - delete sub when caller unsubscribes.
	r.subs = append(r.subs, sub)
	return nil
}

// SubscribersCount returns the number of client subscribed to the node
func (r *RpcClient) SubscribersCount() int32 {
	r.stateMu.RLock()
	defer r.stateMu.RUnlock()
	return int32(len(r.subs))
}

// RPC wrappers

// CallContext implementation
func (r *RpcClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With(
		"method", method,
		"args", args,
	)

	lggr.Debug("RPC call: evmclient.Client#CallContext")
	start := time.Now()
	var err error
	if http != nil {
		err = r.wrapHTTP(http.rpc.CallContext(ctx, result, method, args...))
	} else {
		err = r.wrapWS(ws.rpc.CallContext(ctx, result, method, args...))
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "CallContext")

	return err
}

func (r *RpcClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("nBatchElems", len(b), "batchElems", b)

	lggr.Trace("RPC call: evmclient.Client#BatchCallContext")
	start := time.Now()
	var err error
	if http != nil {
		err = r.wrapHTTP(http.rpc.BatchCallContext(ctx, b))
	} else {
		err = r.wrapWS(ws.rpc.BatchCallContext(ctx, b))
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "BatchCallContext")

	return err
}

func (r *RpcClient) subscribe(ctx context.Context, channel chan<- *evmtypes.Head, args ...interface{}) (commontypes.Subscription, error) {
	ctx, cancel, ws, _ := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("args", args)

	lggr.Debug("RPC call: evmclient.Client#EthSubscribe")
	start := time.Now()
	var sub commontypes.Subscription
	sub, err := ws.rpc.EthSubscribe(ctx, channel, args...)
	if err == nil {
		err = r.registerSub(sub, r.chStopInFlight)
		if err != nil {
			sub.Unsubscribe()
			return nil, err
		}
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "EthSubscribe")

	return sub, r.wrapWS(err)
}

// GethClient wrappers

func (r *RpcClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (receipt *evmtypes.Receipt, err error) {
	err = r.CallContext(ctx, &receipt, "eth_getTransactionReceipt", txHash, false)
	if err != nil {
		return nil, err
	}
	if receipt == nil {
		err = r.wrapRPCClientError(ethereum.NotFound)
		return
	}
	return
}

func (r *RpcClient) TransactionReceiptGeth(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("txHash", txHash)

	lggr.Debug("RPC call: evmclient.Client#TransactionReceipt")

	start := time.Now()
	if http != nil {
		receipt, err = http.geth.TransactionReceipt(ctx, txHash)
		err = r.wrapHTTP(err)
	} else {
		receipt, err = ws.geth.TransactionReceipt(ctx, txHash)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "TransactionReceipt",
		"receipt", receipt,
	)

	return
}
func (r *RpcClient) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("txHash", txHash)

	lggr.Debug("RPC call: evmclient.Client#TransactionByHash")

	start := time.Now()
	if http != nil {
		tx, _, err = http.geth.TransactionByHash(ctx, txHash)
		err = r.wrapHTTP(err)
	} else {
		tx, _, err = ws.geth.TransactionByHash(ctx, txHash)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "TransactionByHash",
		"receipt", tx,
	)

	return
}

func (r *RpcClient) HeaderByNumber(ctx context.Context, number *big.Int) (header *types.Header, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("number", number)

	lggr.Debug("RPC call: evmclient.Client#HeaderByNumber")
	start := time.Now()
	if http != nil {
		header, err = http.geth.HeaderByNumber(ctx, number)
		err = r.wrapHTTP(err)
	} else {
		header, err = ws.geth.HeaderByNumber(ctx, number)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "HeaderByNumber", "header", header)

	return
}

func (r *RpcClient) HeaderByHash(ctx context.Context, hash common.Hash) (header *types.Header, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("hash", hash)

	lggr.Debug("RPC call: evmclient.Client#HeaderByHash")
	start := time.Now()
	if http != nil {
		header, err = http.geth.HeaderByHash(ctx, hash)
		err = r.wrapHTTP(err)
	} else {
		header, err = ws.geth.HeaderByHash(ctx, hash)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "HeaderByHash",
		"header", header,
	)

	return
}

func (r *RpcClient) LatestFinalizedBlock(ctx context.Context) (*evmtypes.Head, error) {
	head, err := r.blockByNumber(ctx, rpc.FinalizedBlockNumber.String())
	if err != nil {
		return nil, err
	}
	return head, nil
}

func (r *RpcClient) BlockByNumber(ctx context.Context, number *big.Int) (head *evmtypes.Head, err error) {
	hex := ToBlockNumArg(number)
	return r.blockByNumber(ctx, hex)
}

func (r *RpcClient) blockByNumber(ctx context.Context, number string) (head *evmtypes.Head, err error) {
	ctx, cancel, chStopInFlight, ws, http := r.acquireQueryCtx(ctx)
	defer cancel()
	const method = "eth_getBlockByNumber"
	args := []interface{}{number, false}
	lggr := r.newRqLggr().With(
		"method", method,
		"args", args,
	)

	lggr.Debug("RPC call: evmclient.Client#CallContext")
	start := time.Now()
	if http != nil {
		err = r.wrapHTTP(http.rpc.CallContext(ctx, &head, method, args...))
	} else {
		err = r.wrapWS(ws.rpc.CallContext(ctx, &head, method, args...))
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "CallContext")
	if err != nil {
		return nil, err
	}
	if head == nil {
		err = r.wrapRPCClientError(ethereum.NotFound)
		return
	}
	head.EVMChainID = ubig.New(r.chainID)

	switch number {
	case rpc.FinalizedBlockNumber.String():
		r.onNewFinalizedHead(ctx, chStopInFlight, head)
	case rpc.LatestBlockNumber.String():
		r.onNewHead(ctx, chStopInFlight, head)
	}

	return
}

func (r *RpcClient) BlockByHash(ctx context.Context, hash common.Hash) (head *evmtypes.Head, err error) {
	err = r.CallContext(ctx, &head, "eth_getBlockByHash", hash.Hex(), false)
	if err != nil {
		return nil, err
	}
	if head == nil {
		err = r.wrapRPCClientError(ethereum.NotFound)
		return
	}
	head.EVMChainID = ubig.New(r.chainID)
	return
}

func (r *RpcClient) BlockByHashGeth(ctx context.Context, hash common.Hash) (block *types.Block, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("hash", hash)

	lggr.Debug("RPC call: evmclient.Client#BlockByHash")
	start := time.Now()
	if http != nil {
		block, err = http.geth.BlockByHash(ctx, hash)
		err = r.wrapHTTP(err)
	} else {
		block, err = ws.geth.BlockByHash(ctx, hash)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "BlockByHash",
		"block", block,
	)

	return
}

func (r *RpcClient) BlockByNumberGeth(ctx context.Context, number *big.Int) (block *types.Block, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("number", number)

	lggr.Debug("RPC call: evmclient.Client#BlockByNumber")
	start := time.Now()
	if http != nil {
		block, err = http.geth.BlockByNumber(ctx, number)
		err = r.wrapHTTP(err)
	} else {
		block, err = ws.geth.BlockByNumber(ctx, number)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "BlockByNumber",
		"block", block,
	)

	return
}

func (r *RpcClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("tx", tx)

	lggr.Debug("RPC call: evmclient.Client#SendTransaction")
	start := time.Now()
	var err error
	if http != nil {
		err = r.wrapHTTP(http.geth.SendTransaction(ctx, tx))
	} else {
		err = r.wrapWS(ws.geth.SendTransaction(ctx, tx))
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "SendTransaction")

	return err
}

func (r *RpcClient) SimulateTransaction(ctx context.Context, tx *types.Transaction) error {
	// Not Implemented
	return pkgerrors.New("SimulateTransaction not implemented")
}

func (r *RpcClient) SendEmptyTransaction(
	ctx context.Context,
	newTxAttempt func(nonce evmtypes.Nonce, feeLimit uint32, fee *assets.Wei, fromAddress common.Address) (attempt any, err error),
	nonce evmtypes.Nonce,
	gasLimit uint32,
	fee *assets.Wei,
	fromAddress common.Address,
) (txhash string, err error) {
	// Not Implemented
	return "", pkgerrors.New("SendEmptyTransaction not implemented")
}

// PendingSequenceAt returns one higher than the highest nonce from both mempool and mined transactions
func (r *RpcClient) PendingSequenceAt(ctx context.Context, account common.Address) (nonce evmtypes.Nonce, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("account", account)

	lggr.Debug("RPC call: evmclient.Client#PendingNonceAt")
	start := time.Now()
	var n uint64
	if http != nil {
		n, err = http.geth.PendingNonceAt(ctx, account)
		nonce = evmtypes.Nonce(int64(n))
		err = r.wrapHTTP(err)
	} else {
		n, err = ws.geth.PendingNonceAt(ctx, account)
		nonce = evmtypes.Nonce(int64(n))
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "PendingNonceAt",
		"nonce", nonce,
	)

	return
}

// SequenceAt is a bit of a misnomer. You might expect it to return the highest
// mined nonce at the given block number, but it actually returns the total
// transaction count which is the highest mined nonce + 1
func (r *RpcClient) SequenceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (nonce evmtypes.Nonce, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("account", account, "blockNumber", blockNumber)

	lggr.Debug("RPC call: evmclient.Client#NonceAt")
	start := time.Now()
	var n uint64
	if http != nil {
		n, err = http.geth.NonceAt(ctx, account, blockNumber)
		nonce = evmtypes.Nonce(int64(n))
		err = r.wrapHTTP(err)
	} else {
		n, err = ws.geth.NonceAt(ctx, account, blockNumber)
		nonce = evmtypes.Nonce(int64(n))
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "NonceAt",
		"nonce", nonce,
	)

	return
}

func (r *RpcClient) PendingCodeAt(ctx context.Context, account common.Address) (code []byte, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("account", account)

	lggr.Debug("RPC call: evmclient.Client#PendingCodeAt")
	start := time.Now()
	if http != nil {
		code, err = http.geth.PendingCodeAt(ctx, account)
		err = r.wrapHTTP(err)
	} else {
		code, err = ws.geth.PendingCodeAt(ctx, account)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "PendingCodeAt",
		"code", code,
	)

	return
}

func (r *RpcClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) (code []byte, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("account", account, "blockNumber", blockNumber)

	lggr.Debug("RPC call: evmclient.Client#CodeAt")
	start := time.Now()
	if http != nil {
		code, err = http.geth.CodeAt(ctx, account, blockNumber)
		err = r.wrapHTTP(err)
	} else {
		code, err = ws.geth.CodeAt(ctx, account, blockNumber)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "CodeAt",
		"code", code,
	)

	return
}

func (r *RpcClient) EstimateGas(ctx context.Context, c interface{}) (gas uint64, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	call := c.(ethereum.CallMsg)
	lggr := r.newRqLggr().With("call", call)

	lggr.Debug("RPC call: evmclient.Client#EstimateGas")
	start := time.Now()
	if http != nil {
		gas, err = http.geth.EstimateGas(ctx, call)
		err = r.wrapHTTP(err)
	} else {
		gas, err = ws.geth.EstimateGas(ctx, call)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "EstimateGas",
		"gas", gas,
	)

	return
}

func (r *RpcClient) SuggestGasPrice(ctx context.Context) (price *big.Int, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr()

	lggr.Debug("RPC call: evmclient.Client#SuggestGasPrice")
	start := time.Now()
	if http != nil {
		price, err = http.geth.SuggestGasPrice(ctx)
		err = r.wrapHTTP(err)
	} else {
		price, err = ws.geth.SuggestGasPrice(ctx)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "SuggestGasPrice",
		"price", price,
	)

	return
}

func (r *RpcClient) CallContract(ctx context.Context, msg interface{}, blockNumber *big.Int) (val []byte, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("callMsg", msg, "blockNumber", blockNumber)
	message := msg.(ethereum.CallMsg)

	lggr.Debug("RPC call: evmclient.Client#CallContract")
	start := time.Now()
	var hex hexutil.Bytes
	if http != nil {
		err = http.rpc.CallContext(ctx, &hex, "eth_call", ToBackwardCompatibleCallArg(message), ToBackwardCompatibleBlockNumArg(blockNumber))
		err = r.wrapHTTP(err)
	} else {
		err = ws.rpc.CallContext(ctx, &hex, "eth_call", ToBackwardCompatibleCallArg(message), ToBackwardCompatibleBlockNumArg(blockNumber))
		err = r.wrapWS(err)
	}
	if err == nil {
		val = hex
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "CallContract",
		"val", val,
	)

	return
}

func (r *RpcClient) PendingCallContract(ctx context.Context, msg interface{}) (val []byte, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("callMsg", msg)
	message := msg.(ethereum.CallMsg)

	lggr.Debug("RPC call: evmclient.Client#PendingCallContract")
	start := time.Now()
	var hex hexutil.Bytes
	if http != nil {
		err = http.rpc.CallContext(ctx, &hex, "eth_call", ToBackwardCompatibleCallArg(message), "pending")
		err = r.wrapHTTP(err)
	} else {
		err = ws.rpc.CallContext(ctx, &hex, "eth_call", ToBackwardCompatibleCallArg(message), "pending")
		err = r.wrapWS(err)
	}
	if err == nil {
		val = hex
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "PendingCallContract",
		"val", val,
	)

	return
}

func (r *RpcClient) LatestBlockHeight(ctx context.Context) (*big.Int, error) {
	var height big.Int
	h, err := r.BlockNumber(ctx)
	return height.SetUint64(h), err
}

func (r *RpcClient) BlockNumber(ctx context.Context) (height uint64, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr()

	lggr.Debug("RPC call: evmclient.Client#BlockNumber")
	start := time.Now()
	if http != nil {
		height, err = http.geth.BlockNumber(ctx)
		err = r.wrapHTTP(err)
	} else {
		height, err = ws.geth.BlockNumber(ctx)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "BlockNumber",
		"height", height,
	)

	return
}

func (r *RpcClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (balance *big.Int, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("account", account.Hex(), "blockNumber", blockNumber)

	lggr.Debug("RPC call: evmclient.Client#BalanceAt")
	start := time.Now()
	if http != nil {
		balance, err = http.geth.BalanceAt(ctx, account, blockNumber)
		err = r.wrapHTTP(err)
	} else {
		balance, err = ws.geth.BalanceAt(ctx, account, blockNumber)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "BalanceAt",
		"balance", balance,
	)

	return
}

// CallArgs represents the data used to call the balance method of a contract.
// "To" is the address of the ERC contract. "Data" is the message sent
// to the contract. "From" is the sender address.
type CallArgs struct {
	From common.Address `json:"from"`
	To   common.Address `json:"to"`
	Data hexutil.Bytes  `json:"data"`
}

// TokenBalance returns the balance of the given address for the token contract address.
func (r *RpcClient) TokenBalance(ctx context.Context, address common.Address, contractAddress common.Address) (*big.Int, error) {
	result := ""
	numLinkBigInt := new(big.Int)
	functionSelector := evmtypes.HexToFunctionSelector(BALANCE_OF_ADDRESS_FUNCTION_SELECTOR) // balanceOf(address)
	data := utils.ConcatBytes(functionSelector.Bytes(), common.LeftPadBytes(address.Bytes(), utils.EVMWordByteLen))
	args := CallArgs{
		To:   contractAddress,
		Data: data,
	}
	err := r.CallContext(ctx, &result, "eth_call", args, "latest")
	if err != nil {
		return numLinkBigInt, err
	}
	if _, ok := numLinkBigInt.SetString(result, 0); !ok {
		return nil, r.wrapRPCClientError(fmt.Errorf("failed to parse int: %s", result))
	}
	return numLinkBigInt, nil
}

// LINKBalance returns the balance of LINK at the given address
func (r *RpcClient) LINKBalance(ctx context.Context, address common.Address, linkAddress common.Address) (*commonassets.Link, error) {
	balance, err := r.TokenBalance(ctx, address, linkAddress)
	if err != nil {
		return commonassets.NewLinkFromJuels(0), err
	}
	return (*commonassets.Link)(balance), nil
}

func (r *RpcClient) FilterEvents(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return r.FilterLogs(ctx, q)
}

func (r *RpcClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) (l []types.Log, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("q", q)

	lggr.Debug("RPC call: evmclient.Client#FilterLogs")
	start := time.Now()
	if http != nil {
		l, err = http.geth.FilterLogs(ctx, q)
		err = r.wrapHTTP(err)
	} else {
		l, err = ws.geth.FilterLogs(ctx, q)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "FilterLogs",
		"log", l,
	)

	return
}

func (r *RpcClient) ClientVersion(ctx context.Context) (version string, err error) {
	err = r.CallContext(ctx, &version, "web3_clientVersion")
	return
}

func (r *RpcClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (_ ethereum.Subscription, err error) {
	ctx, cancel, chStopInFlight, ws, _ := r.acquireQueryCtx(ctx)
	defer cancel()
	lggr := r.newRqLggr().With("q", q)

	lggr.Debug("RPC call: evmclient.Client#SubscribeFilterLogs")
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		r.logResult(lggr, err, duration, r.getRPCDomain(), "SubscribeFilterLogs")
		err = r.wrapWS(err)
	}()
	sub := newSubForwarder(ch, nil, r.wrapRPCClientError)
	err = sub.start(ws.geth.SubscribeFilterLogs(ctx, q, sub.srcCh))
	if err != nil {
		return
	}

	err = r.registerSub(sub, chStopInFlight)
	if err != nil {
		return
	}

	return sub, nil
}

func (r *RpcClient) SuggestGasTipCap(ctx context.Context) (tipCap *big.Int, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr()

	lggr.Debug("RPC call: evmclient.Client#SuggestGasTipCap")
	start := time.Now()
	if http != nil {
		tipCap, err = http.geth.SuggestGasTipCap(ctx)
		err = r.wrapHTTP(err)
	} else {
		tipCap, err = ws.geth.SuggestGasTipCap(ctx)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "SuggestGasTipCap",
		"tipCap", tipCap,
	)

	return
}

// Returns the ChainID according to the geth client. This is useful for functions like verify()
// the common node.
func (r *RpcClient) ChainID(ctx context.Context) (chainID *big.Int, err error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)

	defer cancel()

	if http != nil {
		chainID, err = http.geth.ChainID(ctx)
		err = r.wrapHTTP(err)
	} else {
		chainID, err = ws.geth.ChainID(ctx)
		err = r.wrapWS(err)
	}
	return
}

// newRqLggr generates a new logger with a unique request ID
func (r *RpcClient) newRqLggr() logger.SugaredLogger {
	return r.rpcLog.With("requestID", uuid.New())
}

func (r *RpcClient) wrapRPCClientError(err error) error {
	// simple add msg to the error without adding new stack trace
	return pkgerrors.WithMessage(err, r.rpcClientErrorPrefix())
}

func (r *RpcClient) rpcClientErrorPrefix() string {
	return fmt.Sprintf("RPCClient returned error (%s)", r.name)
}

func wrapCallError(err error, tp string) error {
	if err == nil {
		return nil
	}
	if pkgerrors.Cause(err).Error() == "context deadline exceeded" {
		err = pkgerrors.Wrap(err, "remote node timed out")
	}
	return pkgerrors.Wrapf(err, "%s call failed", tp)
}

func (r *RpcClient) wrapWS(err error) error {
	err = wrapCallError(err, fmt.Sprintf("%s websocket (%s)", r.tier.String(), r.ws.uri.Redacted()))
	return r.wrapRPCClientError(err)
}

func (r *RpcClient) wrapHTTP(err error) error {
	err = wrapCallError(err, fmt.Sprintf("%s http (%s)", r.tier.String(), r.http.uri.Redacted()))
	err = r.wrapRPCClientError(err)
	if err != nil {
		r.rpcLog.Debugw("Call failed", "err", err)
	} else {
		r.rpcLog.Trace("Call succeeded")
	}
	return err
}

// makeLiveQueryCtxAndSafeGetClients wraps makeQueryCtx
func (r *RpcClient) makeLiveQueryCtxAndSafeGetClients(parentCtx context.Context) (ctx context.Context, cancel context.CancelFunc, ws rawclient, http *rawclient) {
	ctx, cancel, _, ws, http = r.acquireQueryCtx(parentCtx)
	return
}

func (r *RpcClient) acquireQueryCtx(parentCtx context.Context) (ctx context.Context, cancel context.CancelFunc,
	chStopInFlight chan struct{}, ws rawclient, http *rawclient) {
	// Need to wrap in mutex because state transition can cancel and replace the
	// context
	r.stateMu.RLock()
	chStopInFlight = r.chStopInFlight
	ws = r.ws
	if r.http != nil {
		cp := *r.http
		http = &cp
	}
	r.stateMu.RUnlock()
	ctx, cancel = makeQueryCtx(parentCtx, chStopInFlight)
	return
}

// makeQueryCtx returns a context that cancels if:
// 1. Passed in ctx cancels
// 2. Passed in channel is closed
// 3. Default timeout is reached (queryTimeout)
func makeQueryCtx(ctx context.Context, ch services.StopChan) (context.Context, context.CancelFunc) {
	var chCancel, timeoutCancel context.CancelFunc
	ctx, chCancel = ch.Ctx(ctx)
	ctx, timeoutCancel = context.WithTimeout(ctx, queryTimeout)
	cancel := func() {
		chCancel()
		timeoutCancel()
	}
	return ctx, cancel
}

func (r *RpcClient) IsSyncing(ctx context.Context) (bool, error) {
	ctx, cancel, ws, http := r.makeLiveQueryCtxAndSafeGetClients(ctx)
	defer cancel()
	lggr := r.newRqLggr()

	lggr.Debug("RPC call: evmclient.Client#SyncProgress")
	var syncProgress *ethereum.SyncProgress
	start := time.Now()
	var err error
	if http != nil {
		syncProgress, err = http.geth.SyncProgress(ctx)
		err = r.wrapHTTP(err)
	} else {
		syncProgress, err = ws.geth.SyncProgress(ctx)
		err = r.wrapWS(err)
	}
	duration := time.Since(start)

	r.logResult(lggr, err, duration, r.getRPCDomain(), "BlockNumber",
		"syncProgress", syncProgress,
	)

	return syncProgress != nil, nil
}

// getChStopInflight provides a convenience helper that mutex wraps a
// read to the chStopInFlight
func (r *RpcClient) getChStopInflight() chan struct{} {
	r.stateMu.RLock()
	defer r.stateMu.RUnlock()
	return r.chStopInFlight
}

func (r *RpcClient) Name() string {
	return r.name
}

func (r *RpcClient) onNewHead(ctx context.Context, requestCh <-chan struct{}, head *evmtypes.Head) {
	if head == nil {
		return
	}

	r.stateMu.Lock()
	defer r.stateMu.Unlock()
	if !commonclient.CtxIsHeathCheckRequest(ctx) {
		r.highestUserObservations.BlockNumber = max(r.highestUserObservations.BlockNumber, head.Number)
		r.highestUserObservations.TotalDifficulty = commonclient.MaxTotalDifficulty(r.highestUserObservations.TotalDifficulty, head.TotalDifficulty)
	}
	select {
	case <-requestCh: // no need to update latestChainInfo, as rpcClient already started new life cycle
		return
	default:
		r.latestChainInfo.BlockNumber = head.Number
		r.latestChainInfo.TotalDifficulty = head.TotalDifficulty
	}
}

func (r *RpcClient) onNewFinalizedHead(ctx context.Context, requestCh <-chan struct{}, head *evmtypes.Head) {
	if head == nil {
		return
	}
	r.stateMu.Lock()
	defer r.stateMu.Unlock()
	if !commonclient.CtxIsHeathCheckRequest(ctx) {
		r.highestUserObservations.FinalizedBlockNumber = max(r.highestUserObservations.FinalizedBlockNumber, head.Number)
	}
	select {
	case <-requestCh: // no need to update latestChainInfo, as rpcClient already started new life cycle
		return
	default:
		r.latestChainInfo.FinalizedBlockNumber = head.Number
	}
}

func (r *RpcClient) GetInterceptedChainInfo() (latest, highestUserObservations commonclient.ChainInfo) {
	r.stateMu.RLock()
	defer r.stateMu.RUnlock()
	return r.latestChainInfo, r.highestUserObservations
}

func ToBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}
