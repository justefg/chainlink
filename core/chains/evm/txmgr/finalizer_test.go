package txmgr_test

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	txmgrcommon "github.com/smartcontractkit/chainlink/v2/common/txmgr"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/headtracker"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
)

func TestFinalizer_MarkTxFinalized(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)
	db := pgtest.NewSqlxDB(t)
	txStore := cltest.NewTestTxStore(t, db)
	ethKeyStore := cltest.NewKeyStore(t, db).Eth()
	feeLimit := uint64(10_000)
	ethClient := testutils.NewEthClientMockWithDefaultChain(t)
	txmClient := txmgr.NewEvmTxmClient(ethClient, nil)
	rpcBatchSize := uint32(1)
	ht := headtracker.NewSimulatedHeadTracker(ethClient, true, 0)

	head := &evmtypes.Head{
		Hash:   utils.NewHash(),
		Number: 100,
		Parent: &evmtypes.Head{
			Hash:        utils.NewHash(),
			Number:      99,
			IsFinalized: true,
		},
	}

	t.Run("returns not finalized for tx with receipt newer than finalized block", func(t *testing.T) {
		finalizer := txmgr.NewEvmFinalizer(logger.Test(t), testutils.FixtureChainID, rpcBatchSize, false, txStore, txmClient, ht)
		servicetest.Run(t, finalizer)

		idempotencyKey := uuid.New().String()
		_, fromAddress := cltest.MustInsertRandomKey(t, ethKeyStore)
		nonce := evmtypes.Nonce(0)
		broadcast := time.Now()
		tx := &txmgr.Tx{
			Sequence:           &nonce,
			IdempotencyKey:     &idempotencyKey,
			FromAddress:        fromAddress,
			EncodedPayload:     []byte{1, 2, 3},
			FeeLimit:           feeLimit,
			State:              txmgrcommon.TxConfirmed,
			BroadcastAt:        &broadcast,
			InitialBroadcastAt: &broadcast,
		}
		attemptHash := insertTxAndAttemptWithIdempotencyKey(t, txStore, tx, idempotencyKey)
		// Insert receipt for unfinalized block num
		mustInsertEthReceipt(t, txStore, head.Number, head.Hash, attemptHash)
		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(head, nil).Once()
		ethClient.On("LatestFinalizedBlock", mock.Anything).Return(head.Parent, nil).Once()
		err := finalizer.ProcessHead(ctx, head)
		require.NoError(t, err)
		tx, err = txStore.FindTxWithIdempotencyKey(ctx, idempotencyKey, testutils.FixtureChainID)
		require.NoError(t, err)
		require.Equal(t, txmgrcommon.TxConfirmed, tx.State)
	})

	t.Run("returns not finalized for tx with receipt re-org'd out", func(t *testing.T) {
		finalizer := txmgr.NewEvmFinalizer(logger.Test(t), testutils.FixtureChainID, rpcBatchSize, false, txStore, txmClient, ht)
		servicetest.Run(t, finalizer)

		idempotencyKey := uuid.New().String()
		_, fromAddress := cltest.MustInsertRandomKey(t, ethKeyStore)
		nonce := evmtypes.Nonce(0)
		broadcast := time.Now()
		tx := &txmgr.Tx{
			Sequence:           &nonce,
			IdempotencyKey:     &idempotencyKey,
			FromAddress:        fromAddress,
			EncodedPayload:     []byte{1, 2, 3},
			FeeLimit:           feeLimit,
			State:              txmgrcommon.TxConfirmed,
			BroadcastAt:        &broadcast,
			InitialBroadcastAt: &broadcast,
		}
		attemptHash := insertTxAndAttemptWithIdempotencyKey(t, txStore, tx, idempotencyKey)
		// Insert receipt for finalized block num
		mustInsertEthReceipt(t, txStore, head.Parent.Number, utils.NewHash(), attemptHash)
		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(head, nil).Once()
		ethClient.On("LatestFinalizedBlock", mock.Anything).Return(head.Parent, nil).Once()
		err := finalizer.ProcessHead(ctx, head)
		require.NoError(t, err)
		tx, err = txStore.FindTxWithIdempotencyKey(ctx, idempotencyKey, testutils.FixtureChainID)
		require.NoError(t, err)
		require.Equal(t, txmgrcommon.TxConfirmed, tx.State)
	})

	t.Run("returns finalized for tx with receipt in a finalized block", func(t *testing.T) {
		finalizer := txmgr.NewEvmFinalizer(logger.Test(t), testutils.FixtureChainID, rpcBatchSize, false, txStore, txmClient, ht)
		servicetest.Run(t, finalizer)

		idempotencyKey := uuid.New().String()
		_, fromAddress := cltest.MustInsertRandomKey(t, ethKeyStore)
		nonce := evmtypes.Nonce(0)
		broadcast := time.Now()
		tx := &txmgr.Tx{
			Sequence:           &nonce,
			IdempotencyKey:     &idempotencyKey,
			FromAddress:        fromAddress,
			EncodedPayload:     []byte{1, 2, 3},
			FeeLimit:           feeLimit,
			State:              txmgrcommon.TxConfirmed,
			BroadcastAt:        &broadcast,
			InitialBroadcastAt: &broadcast,
		}
		attemptHash := insertTxAndAttemptWithIdempotencyKey(t, txStore, tx, idempotencyKey)
		// Insert receipt for finalized block num
		mustInsertEthReceipt(t, txStore, head.Parent.Number, head.Parent.Hash, attemptHash)
		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(head, nil).Once()
		ethClient.On("LatestFinalizedBlock", mock.Anything).Return(head.Parent, nil).Once()
		err := finalizer.ProcessHead(ctx, head)
		require.NoError(t, err)
		tx, err = txStore.FindTxWithIdempotencyKey(ctx, idempotencyKey, testutils.FixtureChainID)
		require.NoError(t, err)
		require.Equal(t, txmgrcommon.TxFinalized, tx.State)
	})

	t.Run("returns finalized for tx with receipt older than block history depth", func(t *testing.T) {
		finalizer := txmgr.NewEvmFinalizer(logger.Test(t), testutils.FixtureChainID, rpcBatchSize, false, txStore, txmClient, ht)
		servicetest.Run(t, finalizer)

		idempotencyKey := uuid.New().String()
		_, fromAddress := cltest.MustInsertRandomKey(t, ethKeyStore)
		nonce := evmtypes.Nonce(0)
		broadcast := time.Now()
		tx := &txmgr.Tx{
			Sequence:           &nonce,
			IdempotencyKey:     &idempotencyKey,
			FromAddress:        fromAddress,
			EncodedPayload:     []byte{1, 2, 3},
			FeeLimit:           feeLimit,
			State:              txmgrcommon.TxConfirmed,
			BroadcastAt:        &broadcast,
			InitialBroadcastAt: &broadcast,
		}
		attemptHash := insertTxAndAttemptWithIdempotencyKey(t, txStore, tx, idempotencyKey)
		// Insert receipt for finalized block num
		receiptBlockHash1 := utils.NewHash()
		mustInsertEthReceipt(t, txStore, head.Parent.Number-2, receiptBlockHash1, attemptHash)
		idempotencyKey = uuid.New().String()
		nonce = evmtypes.Nonce(1)
		tx = &txmgr.Tx{
			Sequence:           &nonce,
			IdempotencyKey:     &idempotencyKey,
			FromAddress:        fromAddress,
			EncodedPayload:     []byte{1, 2, 3},
			FeeLimit:           feeLimit,
			State:              txmgrcommon.TxConfirmed,
			BroadcastAt:        &broadcast,
			InitialBroadcastAt: &broadcast,
		}
		attemptHash = insertTxAndAttemptWithIdempotencyKey(t, txStore, tx, idempotencyKey)
		// Insert receipt for finalized block num
		receiptBlockHash2 := utils.NewHash()
		mustInsertEthReceipt(t, txStore, head.Parent.Number-1, receiptBlockHash2, attemptHash)
		// Separate batch calls will be made for each tx due to RPC batch size set to 1 when finalizer initialized above
		ethClient.On("BatchCallContext", mock.Anything, mock.IsType([]rpc.BatchElem{})).Run(func(args mock.Arguments) {
			rpcElements := args.Get(1).([]rpc.BatchElem)
			require.Equal(t, 1, len(rpcElements))

			require.Equal(t, "eth_getBlockByNumber", rpcElements[0].Method)
			require.Equal(t, false, rpcElements[0].Args[1])

			reqBlockNum := rpcElements[0].Args[0].(string)
			req1BlockNum := hexutil.EncodeBig(big.NewInt(head.Parent.Number - 2))
			req2BlockNum := hexutil.EncodeBig(big.NewInt(head.Parent.Number - 1))
			var headResult evmtypes.Head
			if req1BlockNum == reqBlockNum {
				headResult = evmtypes.Head{Number: head.Parent.Number - 2, Hash: receiptBlockHash1}
			} else if req2BlockNum == reqBlockNum {
				headResult = evmtypes.Head{Number: head.Parent.Number - 1, Hash: receiptBlockHash2}
			} else {
				require.Fail(t, "unrecognized block hash")
			}
			rpcElements[0].Result = &headResult
		}).Return(nil).Twice()
		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(head, nil).Once()
		ethClient.On("LatestFinalizedBlock", mock.Anything).Return(head.Parent, nil).Once()
		err := finalizer.ProcessHead(ctx, head)
		require.NoError(t, err)
		tx, err = txStore.FindTxWithIdempotencyKey(ctx, idempotencyKey, testutils.FixtureChainID)
		require.NoError(t, err)
		require.Equal(t, txmgrcommon.TxFinalized, tx.State)
	})

	t.Run("returns error if failed to retrieve latest head in headtracker", func(t *testing.T) {
		finalizer := txmgr.NewEvmFinalizer(logger.Test(t), testutils.FixtureChainID, rpcBatchSize, false, txStore, txmClient, ht)
		servicetest.Run(t, finalizer)

		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(nil, errors.New("failed to get latest head")).Once()
		err := finalizer.ProcessHead(ctx, head)
		require.Error(t, err)
	})

	t.Run("returns error if failed to calculate latest finalized head in headtracker", func(t *testing.T) {
		finalizer := txmgr.NewEvmFinalizer(logger.Test(t), testutils.FixtureChainID, rpcBatchSize, false, txStore, txmClient, ht)
		servicetest.Run(t, finalizer)

		ethClient.On("HeadByNumber", mock.Anything, mock.Anything).Return(head, nil).Once()
		ethClient.On("LatestFinalizedBlock", mock.Anything).Return(nil, errors.New("failed to calculate latest finalized head")).Once()
		err := finalizer.ProcessHead(ctx, head)
		require.Error(t, err)
	})
}

func insertTxAndAttemptWithIdempotencyKey(t *testing.T, txStore txmgr.TestEvmTxStore, tx *txmgr.Tx, idempotencyKey string) common.Hash {
	ctx := tests.Context(t)
	err := txStore.InsertTx(ctx, tx)
	require.NoError(t, err)
	tx, err = txStore.FindTxWithIdempotencyKey(ctx, idempotencyKey, testutils.FixtureChainID)
	require.NoError(t, err)
	attempt := cltest.NewLegacyEthTxAttempt(t, tx.ID)
	err = txStore.InsertTxAttempt(ctx, &attempt)
	require.NoError(t, err)
	return attempt.Hash
}

// TODO: Update resume pending task run test to work for Finalizer
// func TestEthConfirmer_ResumePendingRuns(t *testing.T) {
// 	t.Parallel()

// 	db := pgtest.NewSqlxDB(t)
// 	config := configtest.NewTestGeneralConfig(t)
// 	txStore := cltest.NewTestTxStore(t, db)

// 	ethKeyStore := cltest.NewKeyStore(t, db).Eth()

// 	_, fromAddress := cltest.MustInsertRandomKey(t, ethKeyStore)

// 	ethClient := testutils.NewEthClientMockWithDefaultChain(t)

// 	evmcfg := evmtest.NewChainScopedConfig(t, config)

// 	head := evmtypes.Head{
// 		Hash:   testutils.NewHash(),
// 		Number: 10,
// 		Parent: &evmtypes.Head{
// 			Hash:   testutils.NewHash(),
// 			Number: 9,
// 			Parent: &evmtypes.Head{
// 				Number: 8,
// 				Hash:   testutils.NewHash(),
// 				Parent: nil,
// 			},
// 		},
// 	}

// 	minConfirmations := int64(2)

// 	pgtest.MustExec(t, db, `SET CONSTRAINTS fk_pipeline_runs_pruning_key DEFERRED`)
// 	pgtest.MustExec(t, db, `SET CONSTRAINTS pipeline_runs_pipeline_spec_id_fkey DEFERRED`)

// 	t.Run("doesn't process task runs that are not suspended (possibly already previously resumed)", func(t *testing.T) {
// 		ec := newEthConfirmer(t, txStore, ethClient, config, evmcfg, ethKeyStore, func(context.Context, uuid.UUID, interface{}, error) error {
// 			t.Fatal("No value expected")
// 			return nil
// 		})

// 		run := cltest.MustInsertPipelineRun(t, db)
// 		tr := cltest.MustInsertUnfinishedPipelineTaskRun(t, db, run.ID)

// 		etx := cltest.MustInsertConfirmedEthTxWithLegacyAttempt(t, txStore, 1, 1, fromAddress)
// 		mustInsertEthReceipt(t, txStore, head.Number-minConfirmations, head.Hash, etx.TxAttempts[0].Hash)
// 		// Setting both signal_callback and callback_completed to TRUE to simulate a completed pipeline task
// 		// It would only be in a state past suspended if the resume callback was called and callback_completed was set to TRUE
// 		pgtest.MustExec(t, db, `UPDATE evm.txes SET pipeline_task_run_id = $1, min_confirmations = $2, signal_callback = TRUE, callback_completed = TRUE WHERE id = $3`, &tr.ID, minConfirmations, etx.ID)

// 		err := ec.ResumePendingTaskRuns(tests.Context(t), &head)
// 		require.NoError(t, err)
// 	})

// 	t.Run("doesn't process task runs where the receipt is younger than minConfirmations", func(t *testing.T) {
// 		ec := newEthConfirmer(t, txStore, ethClient, config, evmcfg, ethKeyStore, func(context.Context, uuid.UUID, interface{}, error) error {
// 			t.Fatal("No value expected")
// 			return nil
// 		})

// 		run := cltest.MustInsertPipelineRun(t, db)
// 		tr := cltest.MustInsertUnfinishedPipelineTaskRun(t, db, run.ID)

// 		etx := cltest.MustInsertConfirmedEthTxWithLegacyAttempt(t, txStore, 2, 1, fromAddress)
// 		mustInsertEthReceipt(t, txStore, head.Number, head.Hash, etx.TxAttempts[0].Hash)

// 		pgtest.MustExec(t, db, `UPDATE evm.txes SET pipeline_task_run_id = $1, min_confirmations = $2, signal_callback = TRUE WHERE id = $3`, &tr.ID, minConfirmations, etx.ID)

// 		err := ec.ResumePendingTaskRuns(tests.Context(t), &head)
// 		require.NoError(t, err)
// 	})

// 	t.Run("processes eth_txes with receipts older than minConfirmations", func(t *testing.T) {
// 		ch := make(chan interface{})
// 		nonce := evmtypes.Nonce(3)
// 		var err error
// 		ec := newEthConfirmer(t, txStore, ethClient, config, evmcfg, ethKeyStore, func(ctx context.Context, id uuid.UUID, value interface{}, thisErr error) error {
// 			err = thisErr
// 			ch <- value
// 			return nil
// 		})

// 		run := cltest.MustInsertPipelineRun(t, db)
// 		tr := cltest.MustInsertUnfinishedPipelineTaskRun(t, db, run.ID)
// 		pgtest.MustExec(t, db, `UPDATE pipeline_runs SET state = 'suspended' WHERE id = $1`, run.ID)

// 		etx := cltest.MustInsertConfirmedEthTxWithLegacyAttempt(t, txStore, int64(nonce), 1, fromAddress)
// 		pgtest.MustExec(t, db, `UPDATE evm.txes SET meta='{"FailOnRevert": true}'`)
// 		receipt := mustInsertEthReceipt(t, txStore, head.Number-minConfirmations, head.Hash, etx.TxAttempts[0].Hash)

// 		pgtest.MustExec(t, db, `UPDATE evm.txes SET pipeline_task_run_id = $1, min_confirmations = $2, signal_callback = TRUE WHERE id = $3`, &tr.ID, minConfirmations, etx.ID)

// 		done := make(chan struct{})
// 		t.Cleanup(func() { <-done })
// 		go func() {
// 			defer close(done)
// 			err2 := ec.ResumePendingTaskRuns(tests.Context(t), &head)
// 			if !assert.NoError(t, err2) {
// 				return
// 			}
// 			// Retrieve Tx to check if callback completed flag was set to true
// 			updateTx, err3 := txStore.FindTxWithSequence(tests.Context(t), fromAddress, nonce)
// 			if assert.NoError(t, err3) {
// 				assert.Equal(t, true, updateTx.CallbackCompleted)
// 			}
// 		}()

// 		select {
// 		case data := <-ch:
// 			assert.NoError(t, err)

// 			require.IsType(t, &evmtypes.Receipt{}, data)
// 			r := data.(*evmtypes.Receipt)
// 			require.Equal(t, receipt.TxHash, r.TxHash)

// 		case <-time.After(time.Second):
// 			t.Fatal("no value received")
// 		}
// 	})

// 	pgtest.MustExec(t, db, `DELETE FROM pipeline_runs`)

// 	t.Run("processes eth_txes with receipt older than minConfirmations that reverted", func(t *testing.T) {
// 		type data struct {
// 			value any
// 			error
// 		}
// 		ch := make(chan data)
// 		nonce := evmtypes.Nonce(4)
// 		ec := newEthConfirmer(t, txStore, ethClient, config, evmcfg, ethKeyStore, func(ctx context.Context, id uuid.UUID, value interface{}, err error) error {
// 			ch <- data{value, err}
// 			return nil
// 		})

// 		run := cltest.MustInsertPipelineRun(t, db)
// 		tr := cltest.MustInsertUnfinishedPipelineTaskRun(t, db, run.ID)
// 		pgtest.MustExec(t, db, `UPDATE pipeline_runs SET state = 'suspended' WHERE id = $1`, run.ID)

// 		etx := cltest.MustInsertConfirmedEthTxWithLegacyAttempt(t, txStore, int64(nonce), 1, fromAddress)
// 		pgtest.MustExec(t, db, `UPDATE evm.txes SET meta='{"FailOnRevert": true}'`)

// 		// receipt is not passed through as a value since it reverted and caused an error
// 		mustInsertRevertedEthReceipt(t, txStore, head.Number-minConfirmations, head.Hash, etx.TxAttempts[0].Hash)

// 		pgtest.MustExec(t, db, `UPDATE evm.txes SET pipeline_task_run_id = $1, min_confirmations = $2, signal_callback = TRUE WHERE id = $3`, &tr.ID, minConfirmations, etx.ID)

// 		done := make(chan struct{})
// 		t.Cleanup(func() { <-done })
// 		go func() {
// 			defer close(done)
// 			err2 := ec.ResumePendingTaskRuns(tests.Context(t), &head)
// 			if !assert.NoError(t, err2) {
// 				return
// 			}
// 			// Retrieve Tx to check if callback completed flag was set to true
// 			updateTx, err3 := txStore.FindTxWithSequence(tests.Context(t), fromAddress, nonce)
// 			if assert.NoError(t, err3) {
// 				assert.Equal(t, true, updateTx.CallbackCompleted)
// 			}
// 		}()

// 		select {
// 		case data := <-ch:
// 			assert.Error(t, data.error)

// 			assert.EqualError(t, data.error, fmt.Sprintf("transaction %s reverted on-chain", etx.TxAttempts[0].Hash.String()))

// 			assert.Nil(t, data.value)

// 		case <-time.After(tests.WaitTimeout(t)):
// 			t.Fatal("no value received")
// 		}
// 	})

// 	t.Run("does not mark callback complete if callback fails", func(t *testing.T) {
// 		nonce := evmtypes.Nonce(5)
// 		ec := newEthConfirmer(t, txStore, ethClient, config, evmcfg, ethKeyStore, func(context.Context, uuid.UUID, interface{}, error) error {
// 			return errors.New("error")
// 		})

// 		run := cltest.MustInsertPipelineRun(t, db)
// 		tr := cltest.MustInsertUnfinishedPipelineTaskRun(t, db, run.ID)

// 		etx := cltest.MustInsertConfirmedEthTxWithLegacyAttempt(t, txStore, int64(nonce), 1, fromAddress)
// 		mustInsertEthReceipt(t, txStore, head.Number-minConfirmations, head.Hash, etx.TxAttempts[0].Hash)
// 		pgtest.MustExec(t, db, `UPDATE evm.txes SET pipeline_task_run_id = $1, min_confirmations = $2, signal_callback = TRUE WHERE id = $3`, &tr.ID, minConfirmations, etx.ID)

// 		err := ec.ResumePendingTaskRuns(tests.Context(t), &head)
// 		require.Error(t, err)

// 		// Retrieve Tx to check if callback completed flag was left unchanged
// 		updateTx, err := txStore.FindTxWithSequence(tests.Context(t), fromAddress, nonce)
// 		require.NoError(t, err)
// 		require.Equal(t, false, updateTx.CallbackCompleted)
// 	})
// }

// func TestEthConfirmer_CheckForReceipts(t *testing.T) {
// 	t.Parallel()

// 	db := pgtest.NewSqlxDB(t)
// 	gconfig, config := newTestChainScopedConfig(t)
// 	txStore := cltest.NewTestTxStore(t, db)

// 	ethClient := testutils.NewEthClientMockWithDefaultChain(t)
// 	ethKeyStore := cltest.NewKeyStore(t, db).Eth()

// 	_, fromAddress := cltest.MustInsertRandomKey(t, ethKeyStore)

// 	ec := newEthConfirmer(t, txStore, ethClient, gconfig, config, ethKeyStore, nil)

// 	nonce := int64(0)
// 	ctx := tests.Context(t)
// 	head := &evmtypes.Head{
// 		Hash:        utils.NewHash(),
// 		Number:      100,
// 		IsFinalized: false,
// 	}

// 	t.Run("only finds eth_txes in unconfirmed state with at least one broadcast attempt", func(t *testing.T) {
// 		mustInsertFatalErrorEthTx(t, txStore, fromAddress)
// 		mustInsertInProgressEthTx(t, txStore, nonce, fromAddress)
// 		nonce++
// 		cltest.MustInsertConfirmedEthTxWithLegacyAttempt(t, txStore, nonce, 1, fromAddress)
// 		nonce++
// 		mustInsertUnconfirmedEthTxWithInsufficientEthAttempt(t, txStore, nonce, fromAddress)
// 		nonce++
// 		mustCreateUnstartedGeneratedTx(t, txStore, fromAddress, config.EVM().ChainID())

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))
// 	})

// 	etx1 := cltest.MustInsertUnconfirmedEthTxWithBroadcastLegacyAttempt(t, txStore, nonce, fromAddress)
// 	nonce++
// 	require.Len(t, etx1.TxAttempts, 1)
// 	attempt1_1 := etx1.TxAttempts[0]
// 	hashAttempt1_1 := attempt1_1.Hash
// 	require.Len(t, attempt1_1.Receipts, 0)

// 	t.Run("fetches receipt for one unconfirmed eth_tx", func(t *testing.T) {
// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		// Transaction not confirmed yet, receipt is nil
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 1 && cltest.BatchElemMatchesParams(b[0], hashAttempt1_1, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			elems[0].Result = &evmtypes.Receipt{}
// 		}).Once()

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		var err error
// 		etx1, err = txStore.FindTxWithAttempts(ctx, etx1.ID)
// 		assert.NoError(t, err)
// 		require.Len(t, etx1.TxAttempts, 1)
// 		attempt1_1 = etx1.TxAttempts[0]
// 		require.NoError(t, err)
// 		require.Len(t, attempt1_1.Receipts, 0)
// 	})

// 	t.Run("saves nothing if returned receipt does not match the attempt", func(t *testing.T) {
// 		txmReceipt := evmtypes.Receipt{
// 			TxHash:           testutils.NewHash(),
// 			BlockHash:        testutils.NewHash(),
// 			BlockNumber:      big.NewInt(42),
// 			TransactionIndex: uint(1),
// 		}

// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		// First transaction confirmed
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 1 && cltest.BatchElemMatchesParams(b[0], hashAttempt1_1, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			*(elems[0].Result.(*evmtypes.Receipt)) = txmReceipt
// 		}).Once()

// 		// No error because it is merely logged
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		etx, err := txStore.FindTxWithAttempts(ctx, etx1.ID)
// 		require.NoError(t, err)
// 		require.Len(t, etx.TxAttempts, 1)

// 		require.Len(t, etx.TxAttempts[0].Receipts, 0)
// 	})

// 	t.Run("saves nothing if query returns error", func(t *testing.T) {
// 		txmReceipt := evmtypes.Receipt{
// 			TxHash:           attempt1_1.Hash,
// 			BlockHash:        testutils.NewHash(),
// 			BlockNumber:      big.NewInt(42),
// 			TransactionIndex: uint(1),
// 		}

// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		// First transaction confirmed
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 1 && cltest.BatchElemMatchesParams(b[0], hashAttempt1_1, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			*(elems[0].Result.(*evmtypes.Receipt)) = txmReceipt
// 			elems[0].Error = errors.New("foo")
// 		}).Once()

// 		// No error because it is merely logged
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		etx, err := txStore.FindTxWithAttempts(ctx, etx1.ID)
// 		require.NoError(t, err)
// 		require.Len(t, etx.TxAttempts, 1)
// 		require.Len(t, etx.TxAttempts[0].Receipts, 0)
// 	})

// 	etx2 := cltest.MustInsertUnconfirmedEthTxWithBroadcastLegacyAttempt(t, txStore, nonce, fromAddress)
// 	nonce++
// 	require.Len(t, etx2.TxAttempts, 1)
// 	attempt2_1 := etx2.TxAttempts[0]
// 	require.Len(t, attempt2_1.Receipts, 0)

// 	t.Run("saves eth_receipt and marks eth_tx as confirmed when geth client returns valid receipt", func(t *testing.T) {
// 		txmReceipt := evmtypes.Receipt{
// 			TxHash:           attempt1_1.Hash,
// 			BlockHash:        testutils.NewHash(),
// 			BlockNumber:      big.NewInt(42),
// 			TransactionIndex: uint(1),
// 			Status:           uint64(1),
// 		}

// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 2 &&
// 				cltest.BatchElemMatchesParams(b[0], attempt1_1.Hash, "eth_getTransactionReceipt") &&
// 				cltest.BatchElemMatchesParams(b[1], attempt2_1.Hash, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			// First transaction confirmed
// 			*(elems[0].Result.(*evmtypes.Receipt)) = txmReceipt
// 			// Second transaction still unconfirmed
// 			elems[1].Result = &evmtypes.Receipt{}
// 		}).Once()

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		// Check that the receipt was saved
// 		etx, err := txStore.FindTxWithAttempts(ctx, etx1.ID)
// 		require.NoError(t, err)

// 		assert.Equal(t, txmgrcommon.TxConfirmed, etx.State)
// 		assert.Len(t, etx.TxAttempts, 1)
// 		attempt1_1 = etx.TxAttempts[0]
// 		require.Len(t, attempt1_1.Receipts, 1)

// 		ethReceipt := attempt1_1.Receipts[0]

// 		assert.Equal(t, txmReceipt.TxHash, ethReceipt.GetTxHash())
// 		assert.Equal(t, txmReceipt.BlockHash, ethReceipt.GetBlockHash())
// 		assert.Equal(t, txmReceipt.BlockNumber.Int64(), ethReceipt.GetBlockNumber().Int64())
// 		assert.Equal(t, txmReceipt.TransactionIndex, ethReceipt.GetTransactionIndex())

// 		receiptJSON, err := json.Marshal(txmReceipt)
// 		require.NoError(t, err)

// 		j, err := json.Marshal(ethReceipt)
// 		require.NoError(t, err)
// 		assert.JSONEq(t, string(receiptJSON), string(j))
// 	})

// 	t.Run("fetches and saves receipts for several attempts in gas price order", func(t *testing.T) {
// 		attempt2_2 := newBroadcastLegacyEthTxAttempt(t, etx2.ID)
// 		attempt2_2.TxFee = gas.EvmFee{Legacy: assets.NewWeiI(10)}

// 		attempt2_3 := newBroadcastLegacyEthTxAttempt(t, etx2.ID)
// 		attempt2_3.TxFee = gas.EvmFee{Legacy: assets.NewWeiI(20)}

// 		// Insert order deliberately reversed to test sorting by gas price
// 		require.NoError(t, txStore.InsertTxAttempt(ctx, &attempt2_3))
// 		require.NoError(t, txStore.InsertTxAttempt(ctx, &attempt2_2))

// 		txmReceipt := evmtypes.Receipt{
// 			TxHash:           attempt2_2.Hash,
// 			BlockHash:        testutils.NewHash(),
// 			BlockNumber:      big.NewInt(42),
// 			TransactionIndex: uint(1),
// 			Status:           uint64(1),
// 		}

// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 3 &&
// 				cltest.BatchElemMatchesParams(b[2], attempt2_1.Hash, "eth_getTransactionReceipt") &&
// 				cltest.BatchElemMatchesParams(b[1], attempt2_2.Hash, "eth_getTransactionReceipt") &&
// 				cltest.BatchElemMatchesParams(b[0], attempt2_3.Hash, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			// Most expensive attempt still unconfirmed
// 			elems[2].Result = &evmtypes.Receipt{}
// 			// Second most expensive attempt is confirmed
// 			*(elems[1].Result.(*evmtypes.Receipt)) = txmReceipt
// 			// Cheapest attempt still unconfirmed
// 			elems[0].Result = &evmtypes.Receipt{}
// 		}).Once()

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		// Check that the state was updated
// 		etx, err := txStore.FindTxWithAttempts(ctx, etx2.ID)
// 		require.NoError(t, err)

// 		require.Equal(t, txmgrcommon.TxConfirmed, etx.State)
// 		require.Len(t, etx.TxAttempts, 3)
// 	})

// 	etx3 := cltest.MustInsertUnconfirmedEthTxWithBroadcastLegacyAttempt(t, txStore, nonce, fromAddress)
// 	attempt3_1 := etx3.TxAttempts[0]
// 	nonce++

// 	t.Run("ignores receipt missing BlockHash that comes from querying parity too early", func(t *testing.T) {
// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		receipt := evmtypes.Receipt{
// 			TxHash: attempt3_1.Hash,
// 			Status: uint64(1),
// 		}
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 1 && cltest.BatchElemMatchesParams(b[0], attempt3_1.Hash, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			*(elems[0].Result.(*evmtypes.Receipt)) = receipt
// 		}).Once()

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		// No receipt, but no error either
// 		etx, err := txStore.FindTxWithAttempts(ctx, etx3.ID)
// 		require.NoError(t, err)

// 		assert.Equal(t, txmgrcommon.TxUnconfirmed, etx.State)
// 		assert.Len(t, etx.TxAttempts, 1)
// 		attempt3_1 = etx.TxAttempts[0]
// 		require.Len(t, attempt3_1.Receipts, 0)
// 	})

// 	t.Run("does not panic if receipt has BlockHash but is missing some other fields somehow", func(t *testing.T) {
// 		// NOTE: This should never happen, but we shouldn't panic regardless
// 		receipt := evmtypes.Receipt{
// 			TxHash:    attempt3_1.Hash,
// 			BlockHash: testutils.NewHash(),
// 			Status:    uint64(1),
// 		}
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 1 && cltest.BatchElemMatchesParams(b[0], attempt3_1.Hash, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			*(elems[0].Result.(*evmtypes.Receipt)) = receipt
// 		}).Once()

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		// No receipt, but no error either
// 		etx, err := txStore.FindTxWithAttempts(ctx, etx3.ID)
// 		require.NoError(t, err)

// 		assert.Equal(t, txmgrcommon.TxUnconfirmed, etx.State)
// 		assert.Len(t, etx.TxAttempts, 1)
// 		attempt3_1 = etx.TxAttempts[0]
// 		require.Len(t, attempt3_1.Receipts, 0)
// 	})
// 	t.Run("handles case where eth_receipt already exists somehow", func(t *testing.T) {
// 		ethReceipt := mustInsertEthReceipt(t, txStore, 42, testutils.NewHash(), attempt3_1.Hash)
// 		txmReceipt := evmtypes.Receipt{
// 			TxHash:           attempt3_1.Hash,
// 			BlockHash:        ethReceipt.BlockHash,
// 			BlockNumber:      big.NewInt(ethReceipt.BlockNumber),
// 			TransactionIndex: ethReceipt.TransactionIndex,
// 			Status:           uint64(1),
// 		}
// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 1 && cltest.BatchElemMatchesParams(b[0], attempt3_1.Hash, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			*(elems[0].Result.(*evmtypes.Receipt)) = txmReceipt
// 		}).Once()

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		// Check that the receipt was unchanged
// 		etx, err := txStore.FindTxWithAttempts(ctx, etx3.ID)
// 		require.NoError(t, err)

// 		assert.Equal(t, txmgrcommon.TxConfirmed, etx.State)
// 		assert.Len(t, etx.TxAttempts, 1)
// 		attempt3_1 = etx.TxAttempts[0]
// 		require.Len(t, attempt3_1.Receipts, 1)

// 		ethReceipt3_1 := attempt3_1.Receipts[0]

// 		assert.Equal(t, txmReceipt.TxHash, ethReceipt3_1.GetTxHash())
// 		assert.Equal(t, txmReceipt.BlockHash, ethReceipt3_1.GetBlockHash())
// 		assert.Equal(t, txmReceipt.BlockNumber.Int64(), ethReceipt3_1.GetBlockNumber().Int64())
// 		assert.Equal(t, txmReceipt.TransactionIndex, ethReceipt3_1.GetTransactionIndex())
// 	})

// 	etx4 := cltest.MustInsertUnconfirmedEthTxWithBroadcastLegacyAttempt(t, txStore, nonce, fromAddress)
// 	attempt4_1 := etx4.TxAttempts[0]
// 	nonce++

// 	t.Run("on receipt fetch marks in_progress eth_tx_attempt as broadcast", func(t *testing.T) {
// 		attempt4_2 := newInProgressLegacyEthTxAttempt(t, etx4.ID)
// 		attempt4_2.TxFee = gas.EvmFee{Legacy: assets.NewWeiI(10)}

// 		require.NoError(t, txStore.InsertTxAttempt(ctx, &attempt4_2))

// 		txmReceipt := evmtypes.Receipt{
// 			TxHash:           attempt4_2.Hash,
// 			BlockHash:        testutils.NewHash(),
// 			BlockNumber:      big.NewInt(42),
// 			TransactionIndex: uint(1),
// 			Status:           uint64(1),
// 		}
// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		// Second attempt is confirmed
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 2 &&
// 				cltest.BatchElemMatchesParams(b[0], attempt4_2.Hash, "eth_getTransactionReceipt") &&
// 				cltest.BatchElemMatchesParams(b[1], attempt4_1.Hash, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			// First attempt still unconfirmed
// 			elems[1].Result = &evmtypes.Receipt{}
// 			// Second attempt is confirmed
// 			*(elems[0].Result.(*evmtypes.Receipt)) = txmReceipt
// 		}).Once()

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		// Check that the state was updated
// 		var err error
// 		etx4, err = txStore.FindTxWithAttempts(ctx, etx4.ID)
// 		require.NoError(t, err)

// 		attempt4_1 = etx4.TxAttempts[1]
// 		attempt4_2 = etx4.TxAttempts[0]

// 		// And the attempts
// 		require.Equal(t, txmgrtypes.TxAttemptBroadcast, attempt4_1.State)
// 		require.Nil(t, attempt4_1.BroadcastBeforeBlockNum)
// 		require.Equal(t, txmgrtypes.TxAttemptBroadcast, attempt4_2.State)
// 		require.Equal(t, int64(42), *attempt4_2.BroadcastBeforeBlockNum)

// 		// Check receipts
// 		require.Len(t, attempt4_1.Receipts, 0)
// 		require.Len(t, attempt4_2.Receipts, 1)
// 	})

// 	etx5 := cltest.MustInsertUnconfirmedEthTxWithBroadcastLegacyAttempt(t, txStore, nonce, fromAddress)
// 	attempt5_1 := etx5.TxAttempts[0]
// 	nonce++

// 	t.Run("simulate on revert", func(t *testing.T) {
// 		txmReceipt := evmtypes.Receipt{
// 			TxHash:           attempt5_1.Hash,
// 			BlockHash:        testutils.NewHash(),
// 			BlockNumber:      big.NewInt(42),
// 			TransactionIndex: uint(1),
// 			Status:           uint64(0),
// 		}
// 		ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 		// First attempt is confirmed and reverted
// 		ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 			return len(b) == 1 &&
// 				cltest.BatchElemMatchesParams(b[0], attempt5_1.Hash, "eth_getTransactionReceipt")
// 		})).Return(nil).Run(func(args mock.Arguments) {
// 			elems := args.Get(1).([]rpc.BatchElem)
// 			// First attempt still unconfirmed
// 			*(elems[0].Result.(*evmtypes.Receipt)) = txmReceipt
// 		}).Once()
// 		data, err := utils.ABIEncode(`[{"type":"uint256"}]`, big.NewInt(10))
// 		require.NoError(t, err)
// 		sig := utils.Keccak256Fixed([]byte(`MyError(uint256)`))
// 		ethClient.On("CallContract", mock.Anything, mock.Anything, mock.Anything).Return(nil, &client.JsonError{
// 			Code:    1,
// 			Message: "reverted",
// 			Data:    utils.ConcatBytes(sig[:4], data),
// 		}).Once()

// 		// Do the thing
// 		require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 		// Check that the state was updated
// 		etx5, err = txStore.FindTxWithAttempts(ctx, etx5.ID)
// 		require.NoError(t, err)

// 		attempt5_1 = etx5.TxAttempts[0]

// 		// And the attempts
// 		require.Equal(t, txmgrtypes.TxAttemptBroadcast, attempt5_1.State)
// 		require.NotNil(t, attempt5_1.BroadcastBeforeBlockNum)
// 		// Check receipts
// 		require.Len(t, attempt5_1.Receipts, 1)
// 	})
// }

// func TestEthConfirmer_CheckForReceipts_batching(t *testing.T) {
// 	t.Parallel()

// 	db := pgtest.NewSqlxDB(t)
// 	cfg := configtest.NewGeneralConfig(t, func(c *chainlink.Config, s *chainlink.Secrets) {
// 		c.EVM[0].RPCDefaultBatchSize = ptr[uint32](2)
// 	})
// 	txStore := cltest.NewTestTxStore(t, db)

// 	ethKeyStore := cltest.NewKeyStore(t, db).Eth()

// 	_, fromAddress := cltest.MustInsertRandomKeyReturningState(t, ethKeyStore)

// 	ethClient := testutils.NewEthClientMockWithDefaultChain(t)

// 	evmcfg := evmtest.NewChainScopedConfig(t, cfg)

// 	ec := newEthConfirmer(t, txStore, ethClient, cfg, evmcfg, ethKeyStore, nil)
// 	ctx := tests.Context(t)

// 	etx := cltest.MustInsertUnconfirmedEthTx(t, txStore, 0, fromAddress)
// 	var attempts []txmgr.TxAttempt
// 	head := &evmtypes.Head{
// 		Hash:        utils.NewHash(),
// 		Number:      42,
// 		IsFinalized: false,
// 	}

// 	// Total of 5 attempts should lead to 3 batched fetches (2, 2, 1)
// 	for i := 0; i < 5; i++ {
// 		attempt := newBroadcastLegacyEthTxAttempt(t, etx.ID, int64(i+2))
// 		require.NoError(t, txStore.InsertTxAttempt(ctx, &attempt))
// 		attempts = append(attempts, attempt)
// 	}

// 	ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)

// 	ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 		return len(b) == 2 &&
// 			cltest.BatchElemMatchesParams(b[0], attempts[4].Hash, "eth_getTransactionReceipt") &&
// 			cltest.BatchElemMatchesParams(b[1], attempts[3].Hash, "eth_getTransactionReceipt")
// 	})).Return(nil).Run(func(args mock.Arguments) {
// 		elems := args.Get(1).([]rpc.BatchElem)
// 		elems[0].Result = &evmtypes.Receipt{}
// 		elems[1].Result = &evmtypes.Receipt{}
// 	}).Once()
// 	ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 		return len(b) == 2 &&
// 			cltest.BatchElemMatchesParams(b[0], attempts[2].Hash, "eth_getTransactionReceipt") &&
// 			cltest.BatchElemMatchesParams(b[1], attempts[1].Hash, "eth_getTransactionReceipt")
// 	})).Return(nil).Run(func(args mock.Arguments) {
// 		elems := args.Get(1).([]rpc.BatchElem)
// 		elems[0].Result = &evmtypes.Receipt{}
// 		elems[1].Result = &evmtypes.Receipt{}
// 	}).Once()
// 	ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 		return len(b) == 1 &&
// 			cltest.BatchElemMatchesParams(b[0], attempts[0].Hash, "eth_getTransactionReceipt")
// 	})).Return(nil).Run(func(args mock.Arguments) {
// 		elems := args.Get(1).([]rpc.BatchElem)
// 		elems[0].Result = &evmtypes.Receipt{}
// 	}).Once()

// 	require.NoError(t, ec.CheckForConfirmation(ctx, head))
// }

// func TestEthConfirmer_CheckForReceipts_HandlesNonFwdTxsWithForwardingEnabled(t *testing.T) {
// 	t.Parallel()

// 	db := pgtest.NewSqlxDB(t)

// 	cfg := configtest.NewGeneralConfig(t, func(c *chainlink.Config, s *chainlink.Secrets) {
// 		c.EVM[0].RPCDefaultBatchSize = ptr[uint32](1)
// 		c.EVM[0].Transactions.ForwardersEnabled = ptr(true)
// 	})

// 	txStore := cltest.NewTestTxStore(t, db)
// 	ethKeyStore := cltest.NewKeyStore(t, db).Eth()
// 	ethClient := testutils.NewEthClientMockWithDefaultChain(t)
// 	evmcfg := evmtest.NewChainScopedConfig(t, cfg)

// 	_, fromAddress := cltest.MustInsertRandomKeyReturningState(t, ethKeyStore)
// 	ec := newEthConfirmer(t, txStore, ethClient, cfg, evmcfg, ethKeyStore, nil)
// 	ctx := tests.Context(t)
// 	head := &evmtypes.Head{
// 		Hash:        utils.NewHash(),
// 		Number:      42,
// 		IsFinalized: false,
// 	}

// 	// tx is not forwarded and doesn't have meta set. EthConfirmer should handle nil meta values
// 	etx := cltest.MustInsertUnconfirmedEthTx(t, txStore, 0, fromAddress)
// 	attempt := newBroadcastLegacyEthTxAttempt(t, etx.ID, 2)
// 	attempt.Tx.Meta = nil
// 	require.NoError(t, txStore.InsertTxAttempt(ctx, &attempt))
// 	dbtx, err := txStore.FindTxWithAttempts(ctx, etx.ID)
// 	require.NoError(t, err)
// 	require.Equal(t, 0, len(dbtx.TxAttempts[0].Receipts))

// 	txmReceipt := evmtypes.Receipt{
// 		TxHash:           attempt.Hash,
// 		BlockHash:        testutils.NewHash(),
// 		BlockNumber:      big.NewInt(42),
// 		TransactionIndex: uint(1),
// 		Status:           uint64(1),
// 	}

// 	ethClient.On("SequenceAt", mock.Anything, mock.Anything, mock.Anything).Return(evmtypes.Nonce(10), nil)
// 	ethClient.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
// 		return len(b) == 1 &&
// 			cltest.BatchElemMatchesParams(b[0], attempt.Hash, "eth_getTransactionReceipt")
// 	})).Return(nil).Run(func(args mock.Arguments) {
// 		elems := args.Get(1).([]rpc.BatchElem)
// 		*(elems[0].Result.(*evmtypes.Receipt)) = txmReceipt // confirmed
// 	}).Once()

// 	require.NoError(t, ec.CheckForConfirmation(ctx, head))

// 	// Check receipt is inserted correctly.
// 	dbtx, err = txStore.FindTxWithAttempts(ctx, etx.ID)
// 	require.NoError(t, err)
// 	require.Equal(t, 1, len(dbtx.TxAttempts[0].Receipts))
// }

func TestFinalizer_FetchReceiptsForConfirmedTx(t *testing.T) {}

// Test checking confirmed tx without receipts that have been broadcasted before the finalized block
func TestFinalizer_CheckConfirmedTxWithoutReceipts(t *testing.T) {}
