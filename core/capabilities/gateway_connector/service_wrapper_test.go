package gatewayconnector

import (
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	chainlink "github.com/smartcontractkit/chainlink/v2/core/services/chainlink/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"
	ksmocks "github.com/smartcontractkit/chainlink/v2/core/services/keystore/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/test-go/testify/mock"
)

// Unit test that creates the ServiceWrapper object and then calls Start() can Close() on it.
// Take inspiration from functions/plugin_test.go and functions/connector_handler_test.go on how to mock the dependencies.
//
// Test valid NodeAddress and an invalid one (i.e. key doesn't exit).

func TestGatewayConnectorServiceWrapper(t *testing.T) {
	t.Parallel()

	logger := logger.TestLogger(t)
	_, addr := testutils.NewPrivateKeyAndAddress(t)

	config, err := chainlink.GeneralConfigOpts{}.New()
	// I don't think this is the right way to get the GatewayConnector in.
	config.Capabilities().GatewayConnector() = chainlink.GatewayConnector{
		ChainIDForNodeKey:         "1",
		NodeAddress:               addr.Hex(),
		DonId:                     "5",
		WSHandshakeTimeoutMillis:  500,
		AuthMinChallengeLen:       0,
		AuthTimestampToleranceSec: 10,
		Gateways:                  []chainlink.ConnectorGateway{{ID: ptr("example_gateway"), URL: ptr("wss://localhost:8081/node")}},
	}
	ethKeystore := ksmocks.NewEth(t)
	ethKeystore.On("EnabledKeysForChain", mock.Anything).Return([]ethkey.KeyV2{{Address: addr}})

	handler := NewGatewayConnectorServiceWrapper(config, ethKeystore, logger)
	require.NoError(t, err)

	t.Cleanup(func() {
		assert.NoError(t, handler.Close())
	})

	t.Run("Start & Stop Success", func(t *testing.T) {
		ctx := testutils.Context(t)

		err := handler.Start(ctx)
		require.NoError(t, err)
		err = handler.Close()
		require.NoError(t, err)
	})

	t.Run("Start Error", func(t *testing.T) {
		// Question, what's the best practices when testing a different configuration, is it 2 configs and then
		// 2 handlers, or does the test do the initialization as in copy from line 27 to 41 into here?
		ctx := testutils.Context(t)
		err := handler.Start(ctx)
		require.NoError(t, err)
	})
}

func ptr[T any](t T) *T { return &t }
