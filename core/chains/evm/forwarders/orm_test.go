package forwarders_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	evmdb "github.com/smartcontractkit/chainlink/v2/core/chains/evm/db"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/forwarders"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	configtest "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/configtest/v2"
	evmtestdb "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/evmtest/db"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type TestORM struct {
	forwarders.ORM
	db *evmdb.ScopedDB
}

func setupORM(t *testing.T) *TestORM {
	t.Helper()

	var (
		cfg  = configtest.NewTestGeneralConfig(t)
		db   = evmtestdb.NewScopedDB(t, cfg.Database())
		lggr = logger.TestLogger(t)
		orm  = forwarders.NewORM(db, lggr, pgtest.NewQConfig(true))
	)

	return &TestORM{ORM: orm, db: db}
}

// Tests the atomicity of cleanup function passed to DeleteForwarder, during DELETE operation
func Test_DeleteForwarder(t *testing.T) {
	t.Parallel()
	orm := setupORM(t)
	addr := testutils.NewAddress()
	chainID := testutils.FixtureChainID

	fwd, err := orm.CreateForwarder(addr, *utils.NewBig(chainID))
	require.NoError(t, err)
	assert.Equal(t, addr, fwd.Address)

	ErrCleaningUp := errors.New("error during cleanup")

	cleanupCalled := 0

	// Cleanup should fail the first time, causing delete to abort.  When cleanup succeeds the second time,
	//  delete should succeed.  Should fail the 3rd and 4th time since the forwarder has already been deleted.
	//  cleanup should only be called the first two times (when DELETE can succeed).
	rets := []error{ErrCleaningUp, nil, nil, ErrCleaningUp}
	expected := []error{ErrCleaningUp, nil, sql.ErrNoRows, sql.ErrNoRows}

	testCleanupFn := func(q pg.Queryer, evmChainID int64, addr common.Address) error {
		require.Less(t, cleanupCalled, len(rets))
		cleanupCalled++
		return rets[cleanupCalled-1]
	}

	for _, expect := range expected {
		err = orm.DeleteForwarder(fwd.ID, testCleanupFn)
		assert.ErrorIs(t, err, expect)
	}
	assert.Equal(t, 2, cleanupCalled)
}
