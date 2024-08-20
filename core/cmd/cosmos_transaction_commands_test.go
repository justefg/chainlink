//go:build integration

package cmd_test

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"

	"github.com/smartcontractkit/chainlink-common/pkg/config"
	cosmosclient "github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/client"
	coscfg "github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/config"
	"github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/denom"
	"github.com/smartcontractkit/chainlink-cosmos/pkg/cosmos/params"

	"github.com/smartcontractkit/chainlink/v2/core/cmd"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/cosmostest"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/cosmoskey"
)

var nativeToken = "cosm"

func TestMain(m *testing.M) {

	params.InitCosmosSdk(
		/* bech32Prefix= */ "wasm",
		/* token= */ nativeToken,
	)

	code := m.Run()
	os.Exit(code)
}

func TestShell_SendCosmosCoins(t *testing.T) {
	ctx := testutils.Context(t)
	// TODO(BCI-978): cleanup once SetupLocalCosmosNode is updated
	chainID := cosmostest.RandomChainID()
	cosmosChain := coscfg.Chain{}
	cosmosChain.SetDefaults()
	accounts, testDir, url := cosmosclient.SetupLocalCosmosNode(t, chainID, *cosmosChain.GasToken)
	require.Greater(t, len(accounts), 1)
	nodes := coscfg.Nodes{
		&coscfg.Node{
			Name:          ptr("random"),
			TendermintURL: config.MustParseURL(url),
		},
	}
	chainConfig := coscfg.TOMLConfig{ChainID: &chainID, Enabled: ptr(true), Chain: cosmosChain, Nodes: nodes}
	app := cosmosStartNewApplication(t, &chainConfig)

	from := accounts[0]
	to := accounts[1]
	require.NoError(t, app.GetKeyStore().Cosmos().Add(ctx, cosmoskey.Raw(from.PrivateKey.Bytes()).Key()))

	require.Eventually(t, func() bool {
		coin, err := cosmosBalance(testDir, url, from.Address.String(), *cosmosChain.GasToken)
		if !assert.NoError(t, err) {
			t.Logf("failed to get balance for %s: %v", from.Address.String(), err)
			return false
		}
		return coin.Sign() > 0
	}, time.Minute, 5*time.Second)

	client, r := app.NewShellAndRenderer()
	cliapp := cli.NewApp()

	for _, tt := range []struct {
		amount string
		expErr string
	}{
		{amount: "0.000001"},
		{amount: "1"},
		{amount: "30.000001"},
		{amount: "1000", expErr: "is too low for this transaction to be executed:"},
		{amount: "0", expErr: "amount must be greater than zero:"},
		{amount: "asdf", expErr: "invalid coin: failed to set decimal string"},
	} {
		tt := tt
		t.Run(tt.amount, func(t *testing.T) {
			startBal, err := cosmosBalance(testDir, url, from.Address.String(), *cosmosChain.GasToken)
			require.NoError(t, err)

			set := flag.NewFlagSet("sendcosmoscoins", 0)
			flagSetApplyFromAction(client.CosmosSendNativeToken, set, "cosmos")

			require.NoError(t, set.Set("id", chainID))
			require.NoError(t, set.Parse([]string{nativeToken, tt.amount, from.Address.String(), to.Address.String()}))

			c := cli.NewContext(cliapp, set, nil)
			err = client.CosmosSendNativeToken(c)
			if tt.expErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expErr)
				return
			}

			// Check CLI output
			require.Greater(t, len(r.Renders), 0)
			renderer := r.Renders[len(r.Renders)-1]
			renderedMsg := renderer.(*cmd.CosmosMsgPresenter)
			require.NotEmpty(t, renderedMsg.ID)
			assert.Equal(t, "unstarted", renderedMsg.State)
			assert.Nil(t, renderedMsg.TxHash)

			// Check balance
			sent, err := denom.ConvertDecCoinToDenom(sdk.NewDecCoinFromDec(nativeToken, sdk.MustNewDecFromStr(tt.amount)), *cosmosChain.GasToken)
			require.NoError(t, err)
			expBal := startBal.Sub(startBal, sent.Amount.BigInt())

			testutils.AssertEventually(t, func() bool {
				endBal, err := cosmosBalance(testDir, url, from.Address.String(), *cosmosChain.GasToken)
				require.NoError(t, err)
				t.Logf("%s <= %s", endBal, expBal)
				return endBal.Cmp(expBal) <= 0
			})
		})
	}
}

func cosmosBalance(homeDir string, tendermintURL string, addr string, denom string) (bal *big.Int, err error) {
	var output []byte
	output, err = exec.Command("wasmd", "--home", homeDir, "query", "bank", "balances", "--denom", denom, addr, "--node", tendermintURL, "--output", "json").Output()
	if err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			err = fmt.Errorf("%v\n%s", err, string(ee.Stderr))
		}
		return
	}
	var raw = struct {
		Amount string `json:"amount"`
		Denom  string `json:"demon"`
	}{}
	err = json.Unmarshal(output, &raw)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal output: %w", err)
		return
	}
	if raw.Denom == denom {
		err = fmt.Errorf("requested denom %s but got %s", denom, raw.Denom)
		return
	}
	if raw.Amount == "" {
		err = errors.New("amount missing")
		return
	}
	bal = new(big.Int)
	bal.SetString(raw.Amount, 10)
	return
}
