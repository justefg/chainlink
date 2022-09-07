// arbgas takes a single URL argument and prints the result of three GetLegacyGas calls to the Arbitrum gas estimator.
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func main() {
	if l := len(os.Args); l != 2 {
		log.Fatal("Expected one URL argument but got", l-1)
	}
	url := os.Args[1]
	lggr, sync := logger.NewLogger()
	defer func() { _ = sync() }()
	lggr.SetLogLevel(zapcore.DebugLevel)

	withEstimator(context.Background(), lggr, url, func(e gas.Estimator) {
		printGetLegacyGas(e, make([]byte, 10), 500_000, assets.GWei(1))
		printGetLegacyGas(e, make([]byte, 10), 500_000, assets.GWei(1), gas.OptForceRefetch)
		printGetLegacyGas(e, make([]byte, 10), max, assets.GWei(1))
	})
}

func printGetLegacyGas(e gas.Estimator, calldata []byte, l2GasLimit uint64, maxGasPrice *big.Int, opts ...gas.Opt) {
	price, limit, err := e.GetLegacyGas(calldata, l2GasLimit, maxGasPrice, opts...)
	if err != nil {
		log.Println("failed to get legacy gas:", err)
		return
	}
	fmt.Println("Price:", (*utils.Wei)(price))
	fmt.Println("Limit:", limit)
}

const max = 50_000_000

func withEstimator(ctx context.Context, lggr logger.Logger, url string, f func(e gas.Estimator)) {
	rc, err := rpc.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	ec := ethclient.NewClient(rc)
	e := gas.NewArbitrumEstimator(lggr, &config{max: max}, rc, ec)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err = e.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer lggr.ErrorIfClosing(e, "ArbitrumEstimator")

	f(e)
}

var _ gas.ArbConfig = &config{}

type config struct {
	max uint64
}

func (c *config) EvmGasLimitMax() uint64 {
	return c.max
}
