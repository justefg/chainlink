package evmtesting

import (
	"math/big"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	clcommontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"

	. "github.com/smartcontractkit/chainlink-common/pkg/types/interfacetests" //nolint common practice to import test mods with .
)

func RunChainComponentsEvmTests[T TestingT[T]](t T, it *EVMChainComponentsInterfaceTester[T]) {
	RunContractReaderEvmTests[T](t, it)
	// Add ChainWriter tests here
}

func RunChainComponentsInLoopEvmTests[T TestingT[T]](t T, it ChainComponentsInterfaceTester[T]) {
	RunContractReaderInLoopTests[T](t, it)
	// Add ChainWriter tests here
}

func RunContractReaderEvmTests[T TestingT[T]](t T, it *EVMChainComponentsInterfaceTester[T]) {
	RunContractReaderInterfaceTests[T](t, it, false)

	t.Run("Dynamically typed topics can be used to filter and have type correct in return", func(t T) {
		it.Setup(t)

		anyString := "foo"
		ctx := it.Helper.Context(t)

		cr := it.GetChainReader(t)
		require.NoError(t, cr.Bind(ctx, it.GetBindings(t)))
		contracts := it.GetBindings(t)
		type DynamicEvent struct {
			Field string
		}
		SubmitTransactionToCW(t, it, "triggerEventWithDynamicTopic", DynamicEvent{Field: anyString}, contracts[0], types.Unconfirmed)

		input := struct{ Field string }{Field: anyString}
		tp := cr.(clcommontypes.ContractTypeProvider)
		output, err := tp.CreateContractType(AnyContractName, triggerWithDynamicTopic, false)
		require.NoError(t, err)
		rOutput := reflect.Indirect(reflect.ValueOf(output))

		require.Eventually(t, func() bool {
			return cr.GetLatestValue(ctx, AnyContractName, triggerWithDynamicTopic, primitives.Unconfirmed, input, output) == nil
		}, it.MaxWaitTimeForEvents(), 100*time.Millisecond)

		assert.Equal(t, &anyString, rOutput.FieldByName("Field").Interface())
		topic, err := abi.MakeTopics([]any{anyString})
		require.NoError(t, err)
		assert.Equal(t, &topic[0][0], rOutput.FieldByName("FieldHash").Interface())
	})

	t.Run("Multiple topics can filter together", func(t T) {
		it.Setup(t)
		ctx := it.Helper.Context(t)
		cr := it.GetChainReader(t)
		require.NoError(t, cr.Bind(ctx, it.GetBindings(t)))

		triggerFourTopics(t, it, int32(1), int32(2), int32(3))
		triggerFourTopics(t, it, int32(2), int32(2), int32(3))
		triggerFourTopics(t, it, int32(1), int32(3), int32(3))
		triggerFourTopics(t, it, int32(1), int32(2), int32(4))

		var latest struct{ Field1, Field2, Field3 int32 }
		params := struct{ Field1, Field2, Field3 int32 }{Field1: 1, Field2: 2, Field3: 3}

		time.Sleep(it.MaxWaitTimeForEvents())

		require.NoError(t, cr.GetLatestValue(ctx, AnyContractName, triggerWithAllTopics, primitives.Unconfirmed, params, &latest))
		assert.Equal(t, int32(1), latest.Field1)
		assert.Equal(t, int32(2), latest.Field2)
		assert.Equal(t, int32(3), latest.Field3)
	})

	t.Run("Filtering can be done on indexed topics that get hashed", func(t T) {
		it.Setup(t)

		cr := it.GetChainReader(t)
		ctx := it.Helper.Context(t)
		require.NoError(t, cr.Bind(ctx, it.GetBindings(t)))

		triggerFourTopicsWithHashed(t, it, "1", [32]uint8{2}, [32]byte{5})
		triggerFourTopicsWithHashed(t, it, "2", [32]uint8{2}, [32]byte{3})
		triggerFourTopicsWithHashed(t, it, "1", [32]uint8{3}, [32]byte{3})

		var latest struct {
			Field3 [32]byte
		}
		params := struct {
			Field1 string
			Field2 [32]uint8
			Field3 [32]byte
		}{Field1: "1", Field2: [32]uint8{2}, Field3: [32]byte{5}}

		time.Sleep(it.MaxWaitTimeForEvents())
		require.NoError(t, cr.GetLatestValue(ctx, AnyContractName, triggerWithAllTopicsWithHashed, primitives.Unconfirmed, params, &latest))
		// only checking Field3 topic makes sense since it isn't hashed, to check other fields we'd have to replicate solidity encoding and hashing
		assert.Equal(t, [32]uint8{5}, latest.Field3)
	})

	t.Run("Bind returns error on missing contract at address", func(t T) {
		it.Setup(t)

		addr := common.BigToAddress(big.NewInt(42))
		reader := it.GetChainReader(t)

		ctx := it.Helper.Context(t)
		err := reader.Bind(ctx, []clcommontypes.BoundContract{{Name: AnyContractName, Address: addr.Hex()}})

		require.ErrorIs(t, err, evm.NoContractExistsError{Address: addr})
	})
}

func RunContractReaderInLoopTests[T TestingT[T]](t T, it ChainComponentsInterfaceTester[T]) {
	RunContractReaderInterfaceTests[T](t, it, false)

	t.Run("Filtering can be done on data words using value comparator", func(t T) {
		it.Setup(t)

		ctx := tests.Context(t)
		cr := it.GetChainReader(t)
		require.NoError(t, cr.Bind(ctx, it.GetBindings(t)))
		contracts := it.GetBindings(t)
		ts1 := CreateTestStruct[T](0, it)
		_ = SubmitTransactionToCW(t, it, MethodTriggeringEvent, ts1, contracts[0], types.Unconfirmed)
		ts2 := CreateTestStruct[T](15, it)
		_ = SubmitTransactionToCW(t, it, MethodTriggeringEvent, ts2, contracts[0], types.Unconfirmed)
		ts3 := CreateTestStruct[T](35, it)
		_ = SubmitTransactionToCW(t, it, MethodTriggeringEvent, ts3, contracts[0], types.Unconfirmed)

		ts := &TestStruct{}
		assert.Eventually(t, func() bool {
			sequences, err := cr.QueryKey(ctx, AnyContractName, query.KeyFilter{Key: EventName, Expressions: []query.Expression{
				query.Comparator("OracleID",
					primitives.ValueComparator{
						Value:    uint8(ts2.OracleID),
						Operator: primitives.Eq,
					}),
			},
			}, query.LimitAndSort{}, ts)
			return err == nil && len(sequences) == 1 && reflect.DeepEqual(&ts2, sequences[0].Data)
		}, it.MaxWaitTimeForEvents(), time.Millisecond*10)
	})
}

func triggerFourTopics[T TestingT[T]](t T, it *EVMChainComponentsInterfaceTester[T], i1, i2, i3 int32) {
	type DynamicEvent struct {
		Field1 int32
		Field2 int32
		Field3 int32
	}
	contracts := it.GetBindings(t)
	SubmitTransactionToCW(t, it, "triggerWithFourTopics", DynamicEvent{Field1: i1, Field2: i2, Field3: i3}, contracts[0], types.Unconfirmed)
}

func triggerFourTopicsWithHashed[T TestingT[T]](t T, it *EVMChainComponentsInterfaceTester[T], i1 string, i2 [32]uint8, i3 [32]byte) {
	type DynamicEvent struct {
		Field1 string
		Field2 [32]uint8
		Field3 [32]byte
	}
	contracts := it.GetBindings(t)
	SubmitTransactionToCW(t, it, "triggerWithFourTopicsWithHashed", DynamicEvent{Field1: i1, Field2: i2, Field3: i3}, contracts[0], types.Unconfirmed)
}
