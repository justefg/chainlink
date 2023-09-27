package headtracker_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	configtest "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/configtest/v2"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/headtracker"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	evmtestdb "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/evmtest/db"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func setupORM(t *testing.T) headtracker.ORM {
	logger := logger.TestLogger(t)
	cfg := configtest.NewGeneralConfig(t, nil)
	evmdb := evmtestdb.NewScopedDB(t, cfg.Database())
	return headtracker.NewORM(evmdb, logger, cfg.Database(), cltest.FixtureChainID)
}
func TestORM_IdempotentInsertHead(t *testing.T) {
	t.Parallel()

	orm := setupORM(t)

	// Returns nil when inserting first head
	head := cltest.Head(0)
	require.NoError(t, orm.IdempotentInsertHead(testutils.Context(t), head))

	// Head is inserted
	foundHead, err := orm.LatestHead(testutils.Context(t))
	require.NoError(t, err)
	assert.Equal(t, head.Hash, foundHead.Hash)

	// Returns nil when inserting same head again
	require.NoError(t, orm.IdempotentInsertHead(testutils.Context(t), head))

	// Head is still inserted
	foundHead, err = orm.LatestHead(testutils.Context(t))
	require.NoError(t, err)
	assert.Equal(t, head.Hash, foundHead.Hash)
}

func TestORM_TrimOldHeads(t *testing.T) {
	t.Parallel()

	orm := setupORM(t)

	for i := 0; i < 10; i++ {
		head := cltest.Head(i)
		require.NoError(t, orm.IdempotentInsertHead(testutils.Context(t), head))
	}

	err := orm.TrimOldHeads(testutils.Context(t), 5)
	require.NoError(t, err)

	heads, err := orm.LatestHeads(testutils.Context(t), 10)
	require.NoError(t, err)

	require.Equal(t, 5, len(heads))
	for i := 0; i < 5; i++ {
		require.LessOrEqual(t, int64(5), heads[i].Number)
	}
}

func TestORM_HeadByHash(t *testing.T) {
	t.Parallel()

	orm := setupORM(t)

	var hash common.Hash
	for i := 0; i < 10; i++ {
		head := cltest.Head(i)
		if i == 5 {
			hash = head.Hash
		}
		require.NoError(t, orm.IdempotentInsertHead(testutils.Context(t), head))
	}

	head, err := orm.HeadByHash(testutils.Context(t), hash)
	require.NoError(t, err)
	require.Equal(t, hash, head.Hash)
	require.Equal(t, int64(5), head.Number)
}

func TestORM_HeadByHash_NotFound(t *testing.T) {
	t.Parallel()

	orm := setupORM(t)

	hash := cltest.Head(123).Hash
	head, err := orm.HeadByHash(testutils.Context(t), hash)

	require.Nil(t, head)
	require.NoError(t, err)
}

func TestORM_LatestHeads_NoRows(t *testing.T) {
	t.Parallel()

	orm := setupORM(t)
	heads, err := orm.LatestHeads(testutils.Context(t), 100)

	require.Zero(t, len(heads))
	require.NoError(t, err)
}
