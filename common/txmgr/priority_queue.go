package txmgr

import (
	"sync"

	feetypes "github.com/smartcontractkit/chainlink/v2/common/fee/types"
	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/common/types"
)

type priorityQueue[
	CHAIN_ID types.ID,
	ADDR, TX_HASH, BLOCK_HASH types.Hashable,
	R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH],
	SEQ types.Sequence,
	FEE feetypes.Fee,
] struct {
	sync.RWMutex
	txs       []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	idToIndex map[int64]int
}

// newPriorityQueue returns a new PriorityQueue instance
func newPriorityQueue[
	CHAIN_ID types.ID,
	ADDR, TX_HASH, BLOCK_HASH types.Hashable,
	R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH],
	SEQ types.Sequence,
	FEE feetypes.Fee,
](maxUnstarted int) *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE] {
	pq := priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]{
		txs:       make([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], 0, maxUnstarted),
		idToIndex: make(map[int64]int),
	}

	return &pq
}

// Close clears the queue
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Close() {
	pq.Lock()
	defer pq.Unlock()

	clear(pq.txs)
	clear(pq.idToIndex)
}

// FindIndexByID returns the index of the transaction with the given ID
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindIndexByID(id int64) int {
	pq.RLock()
	defer pq.RUnlock()

	i, ok := pq.idToIndex[id]
	if !ok {
		return -1
	}
	return i
}

// Peek returns the next transaction to be processed
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Peek() *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE] {
	pq.RLock()
	defer pq.RUnlock()

	if len(pq.txs) == 0 {
		return nil
	}
	return pq.txs[0]
}

// Cap returns the capacity of the queue
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Cap() int {
	pq.RLock()
	defer pq.RUnlock()

	return cap(pq.txs)
}

// Len, Less, Swap, Push, and Pop methods implement the heap interface
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Len() int {
	pq.RLock()
	defer pq.RUnlock()

	return len(pq.txs)
}
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Less(i, j int) bool {
	pq.RLock()
	defer pq.RUnlock()
	// We want Pop to give us the oldest, not newest, transaction based on creation time
	return pq.txs[i].CreatedAt.Before(pq.txs[j].CreatedAt)
}
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Swap(i, j int) {
	pq.Lock()
	defer pq.Unlock()

	pq.txs[i], pq.txs[j] = pq.txs[j], pq.txs[i]
	pq.idToIndex[pq.txs[i].ID] = j
	pq.idToIndex[pq.txs[j].ID] = i
}
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Push(tx any) {
	pq.Lock()
	defer pq.Unlock()

	pq.txs = append(pq.txs, tx.(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]))
}
func (pq *priorityQueue[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Pop() any {
	pq.Lock()
	defer pq.Unlock()

	old := pq.txs
	n := len(old)
	tx := old[n-1]
	old[n-1] = nil // avoid memory leak
	pq.txs = old[0 : n-1]
	return tx
}
