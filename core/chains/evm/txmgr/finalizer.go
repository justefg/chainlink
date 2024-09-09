package txmgr

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
)

var _ Finalizer = (*evmFinalizer)(nil)

var (
	promNumSuccessfulTxs = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "tx_manager_num_successful_transactions",
		Help: "Total number of successful transactions. Note that this can err to be too high since transactions are counted on each confirmation, which can happen multiple times per transaction in the case of re-orgs",
	}, []string{"chainID"})
	promRevertedTxCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "tx_manager_num_tx_reverted",
		Help: "Number of times a transaction reverted on-chain. Note that this can err to be too high since transactions are counted on each confirmation, which can happen multiple times per transaction in the case of re-orgs",
	}, []string{"chainID"})
	promFwdTxCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "tx_manager_fwd_tx_count",
		Help: "The number of forwarded transaction attempts labeled by status",
	}, []string{"chainID", "successful"})
	promTxAttemptCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tx_manager_tx_attempt_count",
		Help: "The number of transaction attempts that are currently being processed by the transaction manager",
	}, []string{"chainID"})
)

// processHeadTimeout represents a sanity limit on how long ProcessHead should take to complete
const processHeadTimeout = 10 * time.Minute

type finalizerTxStore interface {
	FindAttemptsRequiringReceiptFetch(ctx context.Context, chainID *big.Int) (hashes []TxAttempt, err error)
	FindConfirmedTxesReceipts(ctx context.Context, finalizedBlockNum int64, chainID *big.Int) (receipts []*evmtypes.Receipt, err error)
	// Find confirmed txes beyond the minConfirmations param that require callback but have not yet been signaled
	FindTxesPendingCallback(ctx context.Context, blockNum int64, chainID *big.Int) (receiptsPlus []ReceiptPlus, err error)
	PreloadTxes(ctx context.Context, attempts []TxAttempt) error
	SaveFetchedReceipts(ctx context.Context, r []*evmtypes.Receipt, chainID *big.Int) (err error)
	// Update tx to mark that its callback has been signaled
	UpdateTxCallbackCompleted(ctx context.Context, pipelineTaskRunId uuid.UUID, chainId *big.Int) error
	UpdateTxStatesToFinalizedUsingTxHashes(ctx context.Context, txHashes []common.Hash, chainId *big.Int) error
}

type finalizerChainClient interface {
	BatchCallContext(ctx context.Context, elems []rpc.BatchElem) error
	BatchGetReceipts(ctx context.Context, attempts []TxAttempt) (txReceipt []*evmtypes.Receipt, txErr []error, funcErr error)
	CallContract(ctx context.Context, a TxAttempt, blockNumber *big.Int) (rpcErr fmt.Stringer, extractErr error)
}

type finalizerHeadTracker interface {
	LatestAndFinalizedBlock(ctx context.Context) (latest, finalized *evmtypes.Head, err error)
}

type resumeCallback = func(context.Context, uuid.UUID, interface{}, error) error

// Finalizer handles processing new finalized blocks and marking transactions as finalized accordingly in the TXM DB
type evmFinalizer struct {
	services.StateMachine
	lggr              logger.SugaredLogger
	chainId           *big.Int
	rpcBatchSize      int
	forwardersEnabled bool

	txStore     finalizerTxStore
	client      finalizerChainClient
	headTracker finalizerHeadTracker

	mb     *mailbox.Mailbox[*evmtypes.Head]
	stopCh services.StopChan
	wg     sync.WaitGroup

	lastProcessedFinalizedBlockNum int64
	resumeCallback                 resumeCallback
}

func NewEvmFinalizer(
	lggr logger.Logger,
	chainId *big.Int,
	rpcBatchSize uint32,
	forwardersEnabled bool,
	txStore finalizerTxStore,
	client finalizerChainClient,
	headTracker finalizerHeadTracker,
) *evmFinalizer {
	lggr = logger.Named(lggr, "Finalizer")
	return &evmFinalizer{
		lggr:              logger.Sugared(lggr),
		chainId:           chainId,
		rpcBatchSize:      int(rpcBatchSize),
		forwardersEnabled: forwardersEnabled,
		txStore:           txStore,
		client:            client,
		headTracker:       headTracker,
		mb:                mailbox.NewSingle[*evmtypes.Head](),
		resumeCallback:    nil,
	}
}

func (f *evmFinalizer) SetResumeCallback(callback resumeCallback) {
	f.resumeCallback = callback
}

// Start the finalizer
func (f *evmFinalizer) Start(ctx context.Context) error {
	return f.StartOnce("Finalizer", func() error {
		f.lggr.Debugf("started Finalizer with RPC batch size limit: %d", f.rpcBatchSize)
		f.stopCh = make(chan struct{})
		f.wg.Add(1)
		go f.runLoop()
		return nil
	})
}

// Close the finalizer
func (f *evmFinalizer) Close() error {
	return f.StopOnce("Finalizer", func() error {
		f.lggr.Debug("closing Finalizer")
		close(f.stopCh)
		f.wg.Wait()
		return nil
	})
}

func (f *evmFinalizer) Name() string {
	return f.lggr.Name()
}

func (f *evmFinalizer) HealthReport() map[string]error {
	return map[string]error{f.Name(): f.Healthy()}
}

func (f *evmFinalizer) runLoop() {
	defer f.wg.Done()
	ctx, cancel := f.stopCh.NewCtx()
	defer cancel()
	for {
		select {
		case <-f.mb.Notify():
			for {
				if ctx.Err() != nil {
					return
				}
				head, exists := f.mb.Retrieve()
				if !exists {
					break
				}
				if err := f.ProcessHead(ctx, head); err != nil {
					f.lggr.Errorw("Error processing head", "err", err)
					f.SvcErrBuffer.Append(err)
					continue
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (f *evmFinalizer) DeliverLatestHead(head *evmtypes.Head) bool {
	return f.mb.Deliver(head)
}

func (f *evmFinalizer) ProcessHead(ctx context.Context, head *evmtypes.Head) error {
	ctx, cancel := context.WithTimeout(ctx, processHeadTimeout)
	defer cancel()
	err := f.fetchAndStoreReceipts(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch and store receipts for confirmed transactions: %w", err)
	}
	// TODO: Mark transactions broadcasted before the finalized block that we have not gotten a receipt for as fatal using MarkOldTxesMissingReceiptAsErrored. Resume pending task runs with failure
	err = f.ResumePendingTaskRuns(ctx, head)
	if err != nil {
		return fmt.Errorf("failed to resume pending task runs: %w", err)
	}
	_, latestFinalizedHead, err := f.headTracker.LatestAndFinalizedBlock(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve latest finalized head: %w", err)
	}
	return f.processFinalizedHead(ctx, latestFinalizedHead)
}

// processFinalizedHead determines if any confirmed transactions can be marked as finalized by comparing their receipts against the latest finalized block
// Fetches receipts direclty from on-chain so re-org detection is not needed during finalization
func (f *evmFinalizer) processFinalizedHead(ctx context.Context, latestFinalizedHead *evmtypes.Head) error {
	// Cannot determine finality without a finalized head for comparison
	if latestFinalizedHead == nil || !latestFinalizedHead.IsValid() {
		return fmt.Errorf("invalid latestFinalizedHead")
	}
	// Only continue processing if the latestFinalizedHead has not already been processed
	// Helps avoid unnecessary processing on every head if blocks are finalized in batches
	if latestFinalizedHead.BlockNumber() == f.lastProcessedFinalizedBlockNum {
		return nil
	}
	if latestFinalizedHead.BlockNumber() < f.lastProcessedFinalizedBlockNum {
		f.lggr.Errorw("Received finalized block older than one already processed. This should never happen and could be an issue with RPCs.", "lastProcessedFinalizedBlockNum", f.lastProcessedFinalizedBlockNum, "retrievedFinalizedBlockNum", latestFinalizedHead.BlockNumber())
		return nil
	}

	earliestBlockNumInChain := latestFinalizedHead.EarliestHeadInChain().BlockNumber()
	f.lggr.Debugw("processing latest finalized head", "blockNum", latestFinalizedHead.BlockNumber(), "blockHash", latestFinalizedHead.BlockHash(), "earliestBlockNumInChain", earliestBlockNumInChain)

	mark := time.Now()
	// Retrieve all confirmed transactions with receipts older than or equal to the finalized block
	unfinalizedReceipts, err := f.txStore.FindConfirmedTxesReceipts(ctx, latestFinalizedHead.BlockNumber(), f.chainId)
	if err != nil {
		return fmt.Errorf("failed to retrieve receipts for confirmed, unfinalized transactions: %w", err)
	}
	f.lggr.Debugw(fmt.Sprintf("found %d receipts for potential finalized transactions", len(unfinalizedReceipts)), "timeElapsed", time.Since(mark))
	mark = time.Now()

	var finalizedReceipts []*evmtypes.Receipt
	// Group by block hash transactions whose receipts cannot be validated using the cached heads
	blockNumToReceiptsMap := make(map[int64][]*evmtypes.Receipt)
	// Find receipts with block nums older than or equal to the latest finalized block num
	for _, receipt := range unfinalizedReceipts {
		// The tx store query ensures transactions have receipts but leaving this check here for a belts and braces approach
		if receipt.TxHash == utils.EmptyHash || receipt.BlockHash == utils.EmptyHash {
			f.lggr.AssumptionViolationw("invalid receipt found for confirmed transaction", "receipt", receipt)
			continue
		}
		// The tx store query only returns transactions with receipts older than or equal to the finalized block but leaving this check here for a belts and braces approach
		if receipt.BlockNumber.Int64() > latestFinalizedHead.BlockNumber() {
			continue
		}
		// Receipt block num older than earliest head in chain. Validate hash using RPC call later
		if receipt.BlockNumber.Int64() < earliestBlockNumInChain {
			blockNumToReceiptsMap[receipt.BlockNumber.Int64()] = append(blockNumToReceiptsMap[receipt.BlockNumber.Int64()], receipt)
			continue
		}
		blockHashInChain := latestFinalizedHead.HashAtHeight(receipt.BlockNumber.Int64())
		// Receipt block hash does not match the block hash in chain. Transaction has been re-org'd out but DB state has not been updated yet
		if blockHashInChain.String() != receipt.BlockHash.String() {
			// Log error if a transaction is marked as confirmed with a receipt older than the finalized block
			// This scenario could potentially point to a re-org'd transaction the Confirmer has lost track of
			f.lggr.Errorw("found confirmed transaction with re-org'd receipt older than finalized block", "receipt", receipt, "onchainBlockHash", blockHashInChain.String())
			continue
		}
		finalizedReceipts = append(finalizedReceipts, receipt)
	}
	f.lggr.Debugw(fmt.Sprintf("found %d finalized transactions using local block history", len(finalizedReceipts)), "timeElapsed", time.Since(mark))
	mark = time.Now()

	// Check if block hashes exist for receipts on-chain older than the earliest cached head
	// Transactions are grouped by their receipt block hash to avoid repeat requests on the same hash in case transactions were confirmed in the same block
	validatedReceipts := f.batchCheckReceiptHashesOnchain(ctx, blockNumToReceiptsMap)
	finalizedReceipts = append(finalizedReceipts, validatedReceipts...)
	f.lggr.Debugw(fmt.Sprintf("found %d finalized transactions validated against RPC", len(validatedReceipts)), "timeElapsed", time.Since(mark))

	txHashes := f.buildTxHashList(finalizedReceipts)

	mark = time.Now()
	err = f.txStore.UpdateTxStatesToFinalizedUsingTxHashes(ctx, txHashes, f.chainId)
	if err != nil {
		return fmt.Errorf("failed to update transactions as finalized: %w", err)
	}
	f.lggr.Debugw("marked transactions as finalized", "timeElapsed", time.Since(mark))
	return nil
}

func (f *evmFinalizer) batchCheckReceiptHashesOnchain(ctx context.Context, blockNumToReceiptsMap map[int64][]*evmtypes.Receipt) []*evmtypes.Receipt {
	if len(blockNumToReceiptsMap) == 0 {
		return nil
	}
	// Group the RPC batch calls in groups of rpcBatchSize
	var rpcBatchGroups [][]rpc.BatchElem
	var rpcBatch []rpc.BatchElem
	for blockNum := range blockNumToReceiptsMap {
		elem := rpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args: []any{
				hexutil.EncodeBig(big.NewInt(blockNum)),
				false,
			},
			Result: new(evmtypes.Head),
		}
		rpcBatch = append(rpcBatch, elem)
		if len(rpcBatch) >= f.rpcBatchSize {
			rpcBatchGroups = append(rpcBatchGroups, rpcBatch)
			rpcBatch = []rpc.BatchElem{}
		}
	}
	if len(rpcBatch) > 0 {
		rpcBatchGroups = append(rpcBatchGroups, rpcBatch)
	}

	var finalizedReceipts []*evmtypes.Receipt
	for _, rpcBatch := range rpcBatchGroups {
		err := f.client.BatchCallContext(ctx, rpcBatch)
		if err != nil {
			// Continue if batch RPC call failed so other batches can still be considered for finalization
			f.lggr.Errorw("failed to find blocks due to batch call failure", "error", err)
			continue
		}
		for _, req := range rpcBatch {
			if req.Error != nil {
				// Continue if particular RPC call failed so other txs can still be considered for finalization
				f.lggr.Errorw("failed to find block by number", "blockNum", req.Args[0], "error", req.Error)
				continue
			}
			head, ok := req.Result.(*evmtypes.Head)
			if !ok || !head.IsValid() {
				// Continue if particular RPC call yielded a nil block so other txs can still be considered for finalization
				f.lggr.Errorw("retrieved nil head for block number", "blockNum", req.Args[0])
				continue
			}
			receipts := blockNumToReceiptsMap[head.BlockNumber()]
			// Check if transaction receipts match the block hash at the given block num
			// If they do not, the transactions may have been re-org'd out
			// The expectation is for the Confirmer to pick up on these re-orgs and get the transaction included
			for _, receipt := range receipts {
				if receipt.BlockHash.String() == head.BlockHash().String() {
					finalizedReceipts = append(finalizedReceipts, receipt)
				} else {
					// Log error if a transaction is marked as confirmed with a receipt older than the finalized block
					// This scenario could potentially point to a re-org'd transaction the Confirmer has lost track of
					f.lggr.Errorw("found confirmed transaction with re-org'd receipt older than finalized block", "receipt", receipt, "onchainBlockHash", head.BlockHash().String())
				}
			}
		}
	}
	return finalizedReceipts
}

// TODO: Validate attempts broadcastBeforeBlockNum in case all are older or equal to finalized block. If they are, mark them as errored if no receipts found
// TODO: Update fetchAndStoreReceiptsForConfirmedTxes to only fetch receipts for transactions that do not have receipts already stored
func (f *evmFinalizer) fetchAndStoreReceipts(ctx context.Context) error {
	attempts, err := f.txStore.FindAttemptsRequiringReceiptFetch(ctx, f.chainId)
	if err != nil {
		return fmt.Errorf("failed to fetch broadcasted attempts for confirmed transactions: %w", err)
	}
	promTxAttemptCount.WithLabelValues(f.chainId.String()).Set(float64(len(attempts)))
	if len(attempts) == 0 {
		return nil
	}

	f.lggr.Debugw(fmt.Sprintf("Fetching receipts for %v transaction attempts", len(attempts)))

	batchSize := f.rpcBatchSize
	if batchSize == 0 {
		batchSize = len(attempts)
	}
	for i := 0; i < len(attempts); i += batchSize {
		j := i + batchSize
		if j > len(attempts) {
			j = len(attempts)
		}
		batch := attempts[i:j]

		receipts, err := f.batchFetchReceipts(ctx, batch)
		if err != nil {
			return err
		}

		if err := f.txStore.SaveFetchedReceipts(ctx, receipts, f.chainId); err != nil {
			return err
		}
	}

	return nil
}

// Note this function will increment promRevertedTxCount upon receiving
// a reverted transaction receipt. Should only be called with unconfirmed attempts.
func (f *evmFinalizer) batchFetchReceipts(ctx context.Context, attempts []TxAttempt) (receipts []*evmtypes.Receipt, err error) {

	// Metadata is required to determine whether a tx is forwarded or not.
	if f.forwardersEnabled {
		err = f.txStore.PreloadTxes(ctx, attempts)
		if err != nil {
			return nil, fmt.Errorf("Confirmer#batchFetchReceipts error loading txs for attempts: %w", err)
		}
	}

	txReceipts, txErrs, err := f.client.BatchGetReceipts(ctx, attempts)
	if err != nil {
		return nil, err
	}

	for i, receipt := range txReceipts {
		attempt := attempts[i]
		err := txErrs[i]
		if err != nil {
			f.lggr.Error("FetchReceipts failed")
			continue
		}
		ok := f.validateReceipt(ctx, receipt, attempt)
		if !ok {
			continue
		}
		receipts = append(receipts, receipt)
	}

	return
}

func (f *evmFinalizer) validateReceipt(ctx context.Context, receipt *evmtypes.Receipt, attempt TxAttempt) bool {
	l := attempt.Tx.GetLogger(f.lggr).With("txHash", attempt.Hash.String(), "txAttemptID", attempt.ID,
		"txID", attempt.TxID, "nonce", attempt.Tx.Sequence,
	)

	if receipt == nil {
		// NOTE: This should never happen, but it seems safer to check
		// regardless to avoid a potential panic
		l.AssumptionViolation("got nil receipt")
		return false
	}

	if receipt.IsZero() {
		l.Debug("Still waiting for receipt")
		return false
	}

	l = l.With("blockHash", receipt.GetBlockHash().String(), "status", receipt.GetStatus(), "transactionIndex", receipt.GetTransactionIndex())

	if receipt.IsUnmined() {
		l.Debug("Got receipt for transaction but it's still in the mempool and not included in a block yet")
		return false
	}

	l.Debugw("Got receipt for transaction", "blockNumber", receipt.GetBlockNumber(), "feeUsed", receipt.GetFeeUsed())

	if receipt.GetTxHash().String() != attempt.Hash.String() {
		l.Errorf("Invariant violation, expected receipt with hash %s to have same hash as attempt with hash %s", receipt.GetTxHash().String(), attempt.Hash.String())
		return false
	}

	if receipt.GetBlockNumber() == nil {
		l.Error("Invariant violation, receipt was missing block number")
		return false
	}

	if receipt.GetStatus() == 0 {
		if receipt.GetRevertReason() != nil {
			l.Warnw("transaction reverted on-chain", "hash", receipt.GetTxHash(), "revertReason", *receipt.GetRevertReason())
		} else {
			rpcError, errExtract := f.client.CallContract(ctx, attempt, receipt.GetBlockNumber())
			if errExtract == nil {
				l.Warnw("transaction reverted on-chain", "hash", receipt.GetTxHash(), "rpcError", rpcError.String())
			} else {
				l.Warnw("transaction reverted on-chain unable to extract revert reason", "hash", receipt.GetTxHash(), "err", errExtract)
			}
		}
		// This might increment more than once e.g. in case of re-orgs going back and forth we might re-fetch the same receipt
		promRevertedTxCount.WithLabelValues(f.chainId.String()).Add(1)
	} else {
		promNumSuccessfulTxs.WithLabelValues(f.chainId.String()).Add(1)
	}

	// This is only recording forwarded tx that were mined and have a status.
	// Counters are prone to being inaccurate due to re-orgs.
	if f.forwardersEnabled {
		meta, metaErr := attempt.Tx.GetMeta()
		if metaErr == nil && meta != nil && meta.FwdrDestAddress != nil {
			// promFwdTxCount takes two labels, chainId and a boolean of whether a tx was successful or not.
			promFwdTxCount.WithLabelValues(f.chainId.String(), strconv.FormatBool(receipt.GetStatus() != 0)).Add(1)
		}
	}
	return true
}

// ResumePendingTaskRuns issues callbacks to task runs that are pending waiting for receipts
func (f *evmFinalizer) ResumePendingTaskRuns(ctx context.Context, head *evmtypes.Head) error {
	if f.resumeCallback == nil {
		return nil
	}
	receiptsPlus, err := f.txStore.FindTxesPendingCallback(ctx, head.BlockNumber(), f.chainId)

	if err != nil {
		return err
	}

	if len(receiptsPlus) > 0 {
		f.lggr.Debugf("Resuming %d task runs pending receipt", len(receiptsPlus))
	} else {
		f.lggr.Debug("No task runs to resume")
	}
	for _, data := range receiptsPlus {
		var taskErr error
		var output interface{}
		if data.FailOnRevert && data.Receipt.GetStatus() == 0 {
			taskErr = fmt.Errorf("transaction %s reverted on-chain", data.Receipt.GetTxHash())
		} else {
			output = data.Receipt
		}

		f.lggr.Debugw("Callback: resuming tx with receipt", "output", output, "taskErr", taskErr, "pipelineTaskRunID", data.ID)
		if err := f.resumeCallback(ctx, data.ID, output, taskErr); err != nil {
			return fmt.Errorf("failed to resume suspended pipeline run: %w", err)
		}
		// Mark tx as having completed callback
		if err := f.txStore.UpdateTxCallbackCompleted(ctx, data.ID, f.chainId); err != nil {
			return err
		}
	}

	return nil
}

// buildTxHashList builds list of transaction hashes from receipts considered finalized
func (f *evmFinalizer) buildTxHashList(finalizedReceipts []*evmtypes.Receipt) []common.Hash {
	txHashes := make([]common.Hash, len(finalizedReceipts))
	for i, receipt := range finalizedReceipts {
		f.lggr.Debugw("transaction considered finalized",
			"txHash", receipt.TxHash.String(),
			"receiptBlockNum", receipt.BlockNumber.Int64(),
			"receiptBlockHash", receipt.BlockHash.String(),
		)
		txHashes[i] = receipt.TxHash
	}
	return txHashes
}
