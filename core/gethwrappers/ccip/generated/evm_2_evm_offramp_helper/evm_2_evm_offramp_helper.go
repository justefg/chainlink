// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_offramp_helper

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type EVM2EVMOffRampDynamicConfig struct {
	PermissionLessExecutionThresholdSeconds uint32
	Router                                  common.Address
	PriceRegistry                           common.Address
	MaxTokensLength                         uint16
	MaxDataSize                             uint32
}

type EVM2EVMOffRampStaticConfig struct {
	CommitStore         common.Address
	ChainSelector       uint64
	SourceChainSelector uint64
	OnRamp              common.Address
	PrevOffRamp         common.Address
	ArmProxy            common.Address
}

type InternalEVM2EVMMessage struct {
	SourceChainSelector uint64
	Sender              common.Address
	Receiver            common.Address
	SequenceNumber      uint64
	GasLimit            *big.Int
	Strict              bool
	Nonce               uint64
	FeeToken            common.Address
	FeeTokenAmount      *big.Int
	Data                []byte
	TokenAmounts        []ClientEVMTokenAmount
	SourceTokenData     [][]byte
	MessageId           [32]byte
}

type InternalExecutionReport struct {
	Messages          []InternalEVM2EVMMessage
	OffchainTokenData [][][]byte
	Proofs            [][32]byte
	ProofFlagBits     *big.Int
}

type InternalPoolUpdate struct {
	Token common.Address
	Pool  common.Address
}

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type RateLimiterTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

var EVM2EVMOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitStoreAlreadyInUse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenRateLimitError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedSenderWithPreviousRampMessageInflight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IS_SCRIPT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"}],\"name\":\"applyPoolUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"rep\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"manualExecGasLimits\",\"type\":\"uint256[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bitmapIndex\",\"type\":\"uint64\"}],\"name\":\"getExecutionStateBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"}],\"name\":\"releaseOrMintTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"setExecutionStateHelper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"trialExecute\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101806040526014805460ff191660011790553480156200001f57600080fd5b5060405162006eee38038062006eee8339810160408190526200004291620008a6565b838383838033806000816200009e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000d157620000d1816200048d565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080938401819052600380546001600160a01b031916909517600160801b9384021760ff60a01b1916600160a01b90920291909117909355909102909217600455504690528151835114620001815760405162d8548360e71b815260040160405180910390fd5b60608401516001600160a01b03161580620001a4575083516001600160a01b0316155b15620001c3576040516342bcdf7f60e11b815260040160405180910390fd5b83600001516001600160a01b0316634120fccd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000206573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200022c9190620009bd565b6001600160401b03166001146200025657604051636fc2a20760e11b815260040160405180910390fd5b83516001600160a01b0390811660a090815260408601516001600160401b0390811660c05260208701511660e052606086015182166101005260808601518216610140528501511661016052620002cd7fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a62000538565b6101205260005b83518110156200047e576200032e848281518110620002f757620002f7620009db565b6020026020010151848381518110620003145762000314620009db565b6020026020010151600c6200059f60201b9092919060201c565b50620003e3838281518110620003485762000348620009db565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200038e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003b49190620009f1565b848381518110620003c957620003c9620009db565b6020026020010151600f6200059f60201b9092919060201c565b507f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c8482815181106200041a576200041a620009db565b6020026020010151848381518110620004375762000437620009db565b6020026020010151604051620004639291906001600160a01b0392831681529116602082015260400190565b60405180910390a1620004768162000a18565b9050620002d4565b50505050505050505062000a40565b336001600160a01b03821603620004e75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000095565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160c05160e051610100516040516020016200058294939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b6000620005b7846001600160a01b03851684620005bf565b949350505050565b6000620005b784846001600160a01b03851660008281526002840160205260408120829055620005b784846000620005f8838362000601565b90505b92915050565b60008181526001830160205260408120546200064a57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620005fb565b506000620005fb565b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b03811182821017156200068e576200068e62000653565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620006bf57620006bf62000653565b604052919050565b6001600160a01b0381168114620006dd57600080fd5b50565b80516001600160401b0381168114620006f857600080fd5b919050565b60006001600160401b0382111562000719576200071962000653565b5060051b60200190565b600082601f8301126200073557600080fd5b815160206200074e6200074883620006fd565b62000694565b82815260059290921b840181019181810190868411156200076e57600080fd5b8286015b84811015620007965780516200078881620006c7565b835291830191830162000772565b509695505050505050565b600082601f830112620007b357600080fd5b81516020620007c66200074883620006fd565b82815260059290921b84018101918181019086841115620007e657600080fd5b8286015b84811015620007965780516200080081620006c7565b8352918301918301620007ea565b80516001600160801b0381168114620006f857600080fd5b6000606082840312156200083957600080fd5b604051606081016001600160401b03811182821017156200085e576200085e62000653565b8060405250809150825180151581146200087757600080fd5b815262000887602084016200080e565b60208201526200089a604084016200080e565b60408201525092915050565b600080600080848603610160811215620008bf57600080fd5b60c0811215620008ce57600080fd5b50620008d962000669565b8551620008e681620006c7565b8152620008f660208701620006e0565b60208201526200090960408701620006e0565b604082015260608601516200091e81620006c7565b606082015260808601516200093381620006c7565b608082015260a08601516200094881620006c7565b60a082015260c08601519094506001600160401b03808211156200096b57600080fd5b620009798883890162000723565b945060e08701519150808211156200099057600080fd5b506200099f87828801620007a1565b925050620009b286610100870162000826565b905092959194509250565b600060208284031215620009d057600080fd5b620005f882620006e0565b634e487b7160e01b600052603260045260246000fd5b60006020828403121562000a0457600080fd5b815162000a1181620006c7565b9392505050565b60006001820162000a3957634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e051610100516101205161014051610160516163bc62000b3260003960008181610381015281816123190152612cdb015260008181610345015281816118090152818161188b015281816122ef0152818161325101526132d801526000612e94015260008181610309015281816122c801526136f30152600081816102a90152818161227601526136d20152600081816102d9015281816122a0015281816128d0015281816136b10152613eda01526000818161026d015281816122480152612f890152600081816115520152818161159e0152818161194c015261199801526163bc6000f3fe608060405234801561001057600080fd5b50600436106102265760003560e01c80637437ff9f1161012a578063b4069b31116100bd578063d3c7c2c71161008c578063f2fde38b11610071578063f2fde38b146107f4578063f52121a514610807578063f8ccbf471461081a57600080fd5b8063d3c7c2c7146107d9578063d7e2bb50146107e157600080fd5b8063b4069b3114610798578063b5767166146107ab578063c5a1d7f0146107be578063c92b2832146107c657600080fd5b8063856c8247116100f9578063856c82471461071b5780638da5cb5b14610747578063afcb95d714610765578063b1dc65a41461078557600080fd5b80637437ff9f1461060a57806379ba5097146106d557806381ff7048146106dd57806385572ffb1461070d57600080fd5b8063546719cd116101bd578063681fba161161018c5780636c6bd845116101715780636c6bd845146105d1578063704b6c02146105e4578063740f4150146105f757600080fd5b8063681fba161461059c578063693928ae146105b157600080fd5b8063546719cd146104d1578063599f6431146105355780635d86f14114610574578063666cab8d1461058757600080fd5b80632dea00f3116101f95780632dea00f3146104525780633a87ac53146104655780634f9f03fe14610478578063506449721461049957600080fd5b806306285c691461022b578063142a98fc146103d4578063181f5a77146103f45780631ef381741461043d575b600080fd5b6103be6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091526040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815250905090565b6040516103cb9190614942565b60405180910390f35b6103e76103e23660046149d9565b610837565b6040516103cb9190614a60565b6104306040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e312e3000000000000000000000000081525081565b6040516103cb9190614adc565b61045061044b366004614d4a565b6108b2565b005b610450610460366004614e17565b610d71565b610450610473366004614e99565b610d7f565b61048b61048636600461516f565b61118f565b6040516103cb9291906151d3565b6104c36104a73660046149d9565b67ffffffffffffffff1660009081526013602052604090205490565b6040519081526020016103cb565b6104d96111a9565b6040516103cb919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016103cb565b61054f6105823660046151f3565b61125e565b61058f6112c7565b6040516103cb9190615261565b6105a4611336565b6040516103cb9190615274565b6105c46105bf366004615355565b6113ef565b6040516103cb9190615479565b6104506105df366004615675565b611455565b6104506105f23660046151f3565b61145f565b610450610605366004615675565b61154f565b6106c86040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600a5463ffffffff808216835273ffffffffffffffffffffffffffffffffffffffff64010000000090920482166020840152600b549182169383019390935261ffff7401000000000000000000000000000000000000000082041660608301527601000000000000000000000000000000000000000000009004909116608082015290565b6040516103cb9190615730565b6104506116d1565b6007546005546040805163ffffffff808516825264010000000090940490931660208401528201526060016103cb565b610450610226366004615793565b61072e6107293660046151f3565b6117ce565b60405167ffffffffffffffff90911681526020016103cb565b60005473ffffffffffffffffffffffffffffffffffffffff1661054f565b6040805160018152600060208201819052918101919091526060016103cb565b6104506107933660046157ce565b6118f6565b61054f6107a63660046151f3565b611b87565b6104506107b9366004615885565b611c60565b6104c3611c6a565b6104506107d43660046158e7565b611c95565b6105a4611d1a565b61054f6107ef3660046151f3565b611dcf565b6104506108023660046151f3565b611dde565b61045061081536600461516f565b611def565b6014546108279060ff1681565b60405190151581526020016103cb565b600061084560016004615984565b60026108526080856159c6565b67ffffffffffffffff1661086691906159ed565b60136000610875608087615a04565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054901c1660038111156108ac576108ac6149f6565b92915050565b84518460ff16601f821115610928576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f746f6f206d616e79207472616e736d697474657273000000000000000000000060448201526064015b60405180910390fd5b80600003610992576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161091f565b61099a612049565b6109a3856120cc565b60095460005b81811015610a2f5760086000600983815481106109c8576109c8615a2b565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000169055610a2881615a5a565b90506109a9565b50875160005b81811015610c2b5760008a8281518110610a5157610a51615a2b565b6020026020010151905060006002811115610a6e57610a6e6149f6565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260086020526040902054610100900460ff166002811115610aad57610aad6149f6565b14610b14576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161091f565b73ffffffffffffffffffffffffffffffffffffffff8116610b61576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff83168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526008602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610c1157610c116149f6565b02179055509050505080610c2490615a5a565b9050610a35565b508851610c3f9060099060208c01906148ac565b506006805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908b161717905560078054610cc5914691309190600090610c979063ffffffff16615a92565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168d8d8d8d8d8d612379565b6005600001819055506000600760049054906101000a900463ffffffff16905043600760046101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600560000154600760009054906101000a900463ffffffff168e8e8e8e8e8e604051610d5c99989796959493929190615ab5565b60405180910390a15050505050505050505050565b610d7b8282612424565b5050565b610d87612049565b60005b83811015610f86576000858583818110610da657610da6615a2b565b610dbc92602060409092020190810191506151f3565b90506000868684818110610dd257610dd2615a2b565b9050604002016020016020810190610dea91906151f3565b9050610df7600c836124ce565b610e2d576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610e4f600c846124f0565b73ffffffffffffffffffffffffffffffffffffffff1614610e9c576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ea7600c83612512565b50610f228173ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ef6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1a9190615b4b565b600f90612512565b506040805173ffffffffffffffffffffffffffffffffffffffff8085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a1505080610f7f90615a5a565b9050610d8a565b5060005b81811015611188576000838383818110610fa657610fa6615a2b565b610fbc92602060409092020190810191506151f3565b90506000848484818110610fd257610fd2615a2b565b9050604002016020016020810190610fea91906151f3565b905073ffffffffffffffffffffffffffffffffffffffff82161580611023575073ffffffffffffffffffffffffffffffffffffffff8116155b1561105a576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611065600c836124ce565b1561109c576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110a8600c8383612534565b506111248173ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111b9190615b4b565b600f9083612534565b506040805173ffffffffffffffffffffffffffffffffffffffff8085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a150508061118190615a5a565b9050610f8a565b5050505050565b6000606061119d8484612557565b915091505b9250929050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff161515938301939093526004548084166060840152049091166080820152611259906126fa565b905090565b6000808061126d600c856127ac565b91509150816112c0576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161091f565b9392505050565b6060600980548060200260200160405190810160405280929190818152602001828054801561132c57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611301575b5050505050905090565b6060611342600f6127cf565b67ffffffffffffffff81111561135a5761135a614aef565b604051908082528060200260200182016040528015611383578160200160208202803683370190505b50905060005b81518110156113eb57600061139f600f836127da565b509050808383815181106113b5576113b5615a2b565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152506113e481615a5a565b9050611389565b5090565b60606114488989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b925061143991508990508a615b68565b6114438789615b68565b6127f6565b9998505050505050505050565b610d7b8282612cd9565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061149f575060025473ffffffffffffffffffffffffffffffffffffffff163314155b156114d6576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9060200160405180910390a150565b467f0000000000000000000000000000000000000000000000000000000000000000146115da576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015267ffffffffffffffff4616602482015260440161091f565b81515181518114611617576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b818110156116c157600083828151811061163657611636615a2b565b602002602001015190508060001415801561166f5750845180518390811061166057611660615a2b565b60200260200101516080015181105b156116b0576040517f085e39cf000000000000000000000000000000000000000000000000000000008152600481018390526024810182905260440161091f565b506116ba81615a5a565b905061161a565b506116cc8383612cd9565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314611752576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161091f565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b73ffffffffffffffffffffffffffffffffffffffff811660009081526012602052604081205467ffffffffffffffff168015801561184157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1615155b156108ac576040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa1580156118d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c09190615b75565b6119008787613689565b600554883590808214611949576040517f93df584c000000000000000000000000000000000000000000000000000000008152600481018290526024810183905260440161091f565b467f0000000000000000000000000000000000000000000000000000000000000000146119ca576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015246602482015260440161091f565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a13360009081526008602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611a5257611a526149f6565b6002811115611a6357611a636149f6565b9052509050600281602001516002811115611a8057611a806149f6565b148015611ac757506009816000015160ff1681548110611aa257611aa2615a2b565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b611afd576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000611b0b8560206159ed565b611b168860206159ed565b611b228b610144615b92565b611b2c9190615b92565b611b369190615b92565b9050368114611b7a576040517f8e1192e10000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161091f565b5050505050505050505050565b60008080611b96600c856127ac565b9150915081611be9576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161091f565b8073ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c589190615b4b565b949350505050565b610d7b8282613689565b60006112597fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a6136ac565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590611cd5575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15611d0c576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d17600382613779565b50565b6060611d26600c6127cf565b67ffffffffffffffff811115611d3e57611d3e614aef565b604051908082528060200260200182016040528015611d67578160200160208202803683370190505b50905060005b81518110156113eb576000611d83600c836127da565b50905080838381518110611d9957611d99615a2b565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015250611dc881615a5a565b9050611d6d565b6000808061126d600f856127ac565b611de6612049565b611d178161395e565b333014611e28576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611e65565b6040805180820190915260008082526020820152815260200190600190039081611e3e5790505b506101408401515190915015611ed257611ecf8361014001518460200151604051602001611eaf919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528560400151858761016001516127f6565b90505b604083015173ffffffffffffffffffffffffffffffffffffffff163b1580611f3c57506040830151611f3a9073ffffffffffffffffffffffffffffffffffffffff167f85572ffb00000000000000000000000000000000000000000000000000000000613a53565b155b15611f4657505050565b600a546000908190640100000000900473ffffffffffffffffffffffffffffffffffffffff16633cf97983611f7b8786613a6f565b611388886080015189604001516040518563ffffffff1660e01b8152600401611fa79493929190615ba5565b6000604051808303816000875af1158015611fc6573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261200c9190810190615c77565b915091508161118857806040517f0a8d6e8c00000000000000000000000000000000000000000000000000000000815260040161091f9190614adc565b60005473ffffffffffffffffffffffffffffffffffffffff1633146120ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161091f565b565b6000818060200190518101906120e29190615d18565b602081015190915073ffffffffffffffffffffffffffffffffffffffff16612136576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600a805460208085015173ffffffffffffffffffffffffffffffffffffffff908116640100000000027fffffffffffffffff00000000000000000000000000000000000000000000000090931663ffffffff9586161792909217909255604080850151600b80546060808901516080808b0151909916760100000000000000000000000000000000000000000000027fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff61ffff90921674010000000000000000000000000000000000000000027fffffffffffffffffffff00000000000000000000000000000000000000000000909416958816959095179290921791909116929092179055815160c0810183527f00000000000000000000000000000000000000000000000000000000000000008416815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116958201959095527f0000000000000000000000000000000000000000000000000000000000000000909416848301527f00000000000000000000000000000000000000000000000000000000000000008316908401527f00000000000000000000000000000000000000000000000000000000000000008216938301939093527f00000000000000000000000000000000000000000000000000000000000000001660a082015290517f737ef22d3f6615e342ed21c69e06620dbc5c8a261ed7cfb2ce214806b1f76eda9161236d918490615db3565b60405180910390a15050565b6000808a8a8a8a8a8a8a8a8a60405160200161239d99989796959493929190615e82565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b600060026124336080856159c6565b67ffffffffffffffff1661244791906159ed565b90506000601381612459608087615a04565b67ffffffffffffffff16815260208101919091526040016000205490508161248360016004615984565b901b19168183600381111561249a5761249a6149f6565b901b1780601360006124ad608088615a04565b67ffffffffffffffff16815260208101919091526040016000205550505050565b60006112c08373ffffffffffffffffffffffffffffffffffffffff8416613b1f565b60006112c08373ffffffffffffffffffffffffffffffffffffffff8416613b2b565b60006112c08373ffffffffffffffffffffffffffffffffffffffff8416613b37565b6000611c588473ffffffffffffffffffffffffffffffffffffffff851684613b43565b6040517ff52121a5000000000000000000000000000000000000000000000000000000008152600090606090309063f52121a59061259b9087908790600401615f6c565b600060405180830381600087803b1580156125b557600080fd5b505af19250505080156125c6575060015b6126df573d8080156125f4576040519150601f19603f3d011682016040523d82523d6000602084013e6125f9565b606091505b50612603816160f6565b7fffffffff00000000000000000000000000000000000000000000000000000000167f0a8d6e8c00000000000000000000000000000000000000000000000000000000148061269b5750612656816160f6565b7fffffffff00000000000000000000000000000000000000000000000000000000167fe1cd550900000000000000000000000000000000000000000000000000000000145b156126ab576003925090506111a2565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161091f9190614adc565b50506040805160208101909152600081526002909250929050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261278882606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff164261276c9190615984565b85608001516fffffffffffffffffffffffffffffffff16613b66565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b60008061119d8473ffffffffffffffffffffffffffffffffffffffff8516613b85565b60006108ac82613b94565b60008080806127e98686613b9f565b9097909650945050505050565b60606000865167ffffffffffffffff81111561281457612814614aef565b60405190808252806020026020018201604052801561285957816020015b60408051808201909152600080825260208201528152602001906001900390816128325790505b50905060005b8751811015612ca857600061289089838151811061287f5761287f615a2b565b60200260200101516000015161125e565b90508073ffffffffffffffffffffffffffffffffffffffff16638627fad689898c86815181106128c2576128c2615a2b565b6020026020010151602001517f00000000000000000000000000000000000000000000000000000000000000008b888151811061290157612901615a2b565b60200260200101518b898151811061291b5761291b615a2b565b6020026020010151604051602001612934929190616146565b6040516020818303038152906040526040518663ffffffff1660e01b815260040161296395949392919061616b565b600060405180830381600087803b15801561297d57600080fd5b505af192505050801561298e575060015b612bb1573d8080156129bc576040519150601f19603f3d011682016040523d82523d6000602084013e6129c1565b606091505b5060006129cd826160f6565b90507f9725942a000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082161480612a6057507ff94ebcd1000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80612aac57507f15279c08000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80612af857507f1a76572a000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80612b4457507fd0c8d23a000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b15612b7d57816040517f30dabb5900000000000000000000000000000000000000000000000000000000815260040161091f9190614adc565b816040517fe1cd550900000000000000000000000000000000000000000000000000000000815260040161091f9190614adc565b8073ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612bfc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c209190615b4b565b838381518110612c3257612c32615a2b565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690528851899083908110612c6b57612c6b615a2b565b602002602001015160200151838381518110612c8957612c89615a2b565b602090810291909101810151015250612ca181615a5a565b905061285f565b50600b54612ccd90829073ffffffffffffffffffffffffffffffffffffffff16613bae565b90505b95945050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612d44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d6891906161ce565b15612d9f576040517fc148371500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151516000819003612ddc576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260200151518114612e1a576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115612e3557612e35614aef565b604051908082528060200260200182016040528015612e5e578160200160208202803683370190505b50905060005b82811015612f3e57600085600001518281518110612e8457612e84615a2b565b60200260200101519050612eb8817f0000000000000000000000000000000000000000000000000000000000000000613d8e565b838381518110612eca57612eca615a2b565b602002602001018181525050806101800151838381518110612eee57612eee615a2b565b602002602001015114612f2d576040517f7185cf6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50612f3781615a5a565b9050612e64565b50604080850151606086015191517f3204887500000000000000000000000000000000000000000000000000000000815260009273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692633204887592612fbf9287929160040161621b565b602060405180830381865afa158015612fdc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130009190616251565b90508060000361303c576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8351151560005b848110156136805760008760000151828151811061306357613063615a2b565b60200260200101519050600061307c8260600151610837565b90506000816003811115613092576130926149f6565b14806130af575060038160038111156130ad576130ad6149f6565b145b6130f75760608201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161091f565b83156131b457600a5460009063ffffffff166131138742615984565b119050808061313357506003826003811115613131576131316149f6565b145b613169576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b88848151811061317b5761317b615a2b565b60200260200101516000146131ae5788848151811061319c5761319c615a2b565b60200260200101518360800181815250505b50613211565b60008160038111156131c8576131c86149f6565b146132115760608201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161091f565b60208083015173ffffffffffffffffffffffffffffffffffffffff1660009081526012909152604090205467ffffffffffffffff168015801561328957507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1615155b1561342c5760208301516040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201527f00000000000000000000000000000000000000000000000000000000000000009091169063856c824790602401602060405180830381865afa158015613321573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133459190615b75565b60c084015190915067ffffffffffffffff1661336282600161626a565b67ffffffffffffffff16146133cf57826020015173ffffffffffffffffffffffffffffffffffffffff168360c0015167ffffffffffffffff167fe44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5d60405160405180910390a3505050613670565b60208381015173ffffffffffffffffffffffffffffffffffffffff16600090815260129091526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83161790555b6000826003811115613440576134406149f6565b036134cc5760c083015167ffffffffffffffff1661345f82600161626a565b67ffffffffffffffff16146134cc57826020015173ffffffffffffffffffffffffffffffffffffffff168360c0015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a3505050613670565b60008a6020015185815181106134e4576134e4615a2b565b602002602001015190506134f9848251613ed8565b61350884606001516001612424565b6000806135158684612557565b91509150613527866060015183612424565b600382600381111561353b5761353b6149f6565b1415801561355b57506002826003811115613558576135586149f6565b14155b1561359a578560600151826040517f9e26160300000000000000000000000000000000000000000000000000000000815260040161091f92919061628b565b60008560038111156135ae576135ae6149f6565b0361361b5760208087015173ffffffffffffffffffffffffffffffffffffffff166000908152601290915260408120805467ffffffffffffffff16916135f3836162a9565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b856101800151866060015167ffffffffffffffff167fd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef6584846040516136619291906151d3565b60405180910390a35050505050505b61367981615a5a565b9050613043565b50505050505050565b610d7b613698828401846162c6565b604080516000815260208101909152612cd9565b6000817f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060405160200161375c949392919093845267ffffffffffffffff92831660208501529116604083015273ffffffffffffffffffffffffffffffffffffffff16606082015260800190565b604051602081830303815290604052805190602001209050919050565b81546000906137a290700100000000000000000000000000000000900463ffffffff1642615984565b9050801561384457600183015483546137ea916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416613b66565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b6020820151835461386a916fffffffffffffffffffffffffffffffff90811691166140ae565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19906139519084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b60405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036139dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161091f565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613a5e836140c4565b80156112c057506112c08383614128565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461018001518152602001846000015167ffffffffffffffff1681526020018460200151604051602001613af4919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001846101200151815260200183815250905092915050565b60006112c083836141f7565b60006112c08383614203565b60006112c0838361428d565b6000611c58848473ffffffffffffffffffffffffffffffffffffffff85166142aa565b6000612cd085613b7684866159ed565b613b809087615b92565b6140ae565b60008080806127e986866142c7565b60006108ac82614301565b60008080806127e9868661430c565b81516000805b82811015613d7a5760008473ffffffffffffffffffffffffffffffffffffffff1663d02641a0878481518110613bec57613bec615a2b565b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016040805180830381865afa158015613c60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c8491906162fb565b51905077ffffffffffffffffffffffffffffffffffffffffffffffff8116600003613d1257858281518110613cbb57613cbb615a2b565b6020908102919091010151516040517f9a655f7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015260240161091f565b613d5c868381518110613d2757613d27615a2b565b6020026020010151602001518277ffffffffffffffffffffffffffffffffffffffffffffffff1661433790919063ffffffff16565b613d669084615b92565b92505080613d7390615a5a565b9050613bb4565b50613d886003826000614370565b50505050565b60008060001b8284606001518560c001518660200151876040015188610120015180519060200120896101400151604051602001613dcc9190615479565b604051602081830303815290604052805190602001208a6101600151604051602001613df89190616359565b604051602081830303815290604052805190602001208b608001518c60a001518d60e001518e6101000151604051602001613eba9d9c9b9a999897969594939291909c8d5260208d019b909b5267ffffffffffffffff998a1660408d01529790981660608b015273ffffffffffffffffffffffffffffffffffffffff95861660808b015293851660a08a015260c089019290925260e0880152610100870152610120860152911515610140850152166101608301526101808201526101a00190565b60405160208183030381529060405280519060200120905092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16826000015167ffffffffffffffff1614613f585781516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161091f565b600b54610140830151517401000000000000000000000000000000000000000090910461ffff161015613fc95760608201516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161091f565b80826101400151511461401a5760608201516040517f8808f8e700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161091f565b600b546101208301515176010000000000000000000000000000000000000000000090910463ffffffff161015610d7b57600b54610120830151516040517f8693378900000000000000000000000000000000000000000000000000000000815276010000000000000000000000000000000000000000000090920463ffffffff166004830152602482015260440161091f565b60008183106140bd57816112c0565b5090919050565b60006140f0827f01ffc9a700000000000000000000000000000000000000000000000000000000614128565b80156108ac5750614121827fffffffff00000000000000000000000000000000000000000000000000000000614128565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156141e0575060208210155b80156141ec5750600081115b979650505050505050565b60006112c083836146f3565b600081815260028301602052604081205480151580614227575061422784846141f7565b6112c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b65790000604482015260640161091f565b600081815260028301602052604081208190556112c0838361470b565b60008281526002840160205260408120829055611c588484614717565b60008181526002830160205260408120548190806142f6576142e985856141f7565b9250600091506111a29050565b6001925090506111a2565b60006108ac82614723565b6000808061431a858561472d565b600081815260029690960160205260409095205494959350505050565b6000670de0b6b3a76400006143668377ffffffffffffffffffffffffffffffffffffffffffffffff86166159ed565b6112c0919061636c565b825474010000000000000000000000000000000000000000900460ff161580614397575081155b156143a157505050565b825460018401546fffffffffffffffffffffffffffffffff808316929116906000906143e790700100000000000000000000000000000000900463ffffffff1642615984565b905080156144a75781831115614429576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546144639083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16613b66565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b8482101561455e5773ffffffffffffffffffffffffffffffffffffffff8416614506576040517ff94ebcd1000000000000000000000000000000000000000000000000000000008152600481018390526024810186905260440161091f565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff8516604482015260640161091f565b848310156146715760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff169060009082906145a29082615984565b6145ac878a615984565b6145b69190615b92565b6145c0919061636c565b905073ffffffffffffffffffffffffffffffffffffffff8616614619576040517f15279c08000000000000000000000000000000000000000000000000000000008152600481018290526024810186905260440161091f565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff8716604482015260640161091f565b61467b8584615984565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b600081815260018301602052604081205415156112c0565b60006112c08383614739565b60006112c08383614833565b60006108ac825490565b60006112c08383614882565b6000818152600183016020526040812054801561482257600061475d600183615984565b855490915060009061477190600190615984565b90508181146147d657600086600001828154811061479157614791615a2b565b90600052602060002001549050808760000184815481106147b4576147b4615a2b565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806147e7576147e7616380565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506108ac565b60009150506108ac565b5092915050565b600081815260018301602052604081205461487a575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556108ac565b5060006108ac565b600082600001828154811061489957614899615a2b565b9060005260206000200154905092915050565b828054828255906000526020600020908101928215614926579160200282015b8281111561492657825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161782556020909201916001909101906148cc565b506113eb9291505b808211156113eb576000815560010161492e565b60c081016108ac828473ffffffffffffffffffffffffffffffffffffffff808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015250508060608301511660608401528060808301511660808401528060a08301511660a0840152505050565b67ffffffffffffffff81168114611d1757600080fd5b80356149d4816149b3565b919050565b6000602082840312156149eb57600080fd5b81356112c0816149b3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110614a5c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016108ac8284614a25565b60005b83811015614a89578181015183820152602001614a71565b50506000910152565b60008151808452614aaa816020860160208601614a6e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006112c06020830184614a92565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614b4157614b41614aef565b60405290565b6040516101a0810167ffffffffffffffff81118282101715614b4157614b41614aef565b6040516080810167ffffffffffffffff81118282101715614b4157614b41614aef565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614bd557614bd5614aef565b604052919050565b600067ffffffffffffffff821115614bf757614bf7614aef565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81168114611d1757600080fd5b80356149d481614c01565b600082601f830112614c3f57600080fd5b81356020614c54614c4f83614bdd565b614b8e565b82815260059290921b84018101918181019086841115614c7357600080fd5b8286015b84811015614c97578035614c8a81614c01565b8352918301918301614c77565b509695505050505050565b803560ff811681146149d457600080fd5b600067ffffffffffffffff821115614ccd57614ccd614aef565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112614d0a57600080fd5b8135614d18614c4f82614cb3565b818152846020838601011115614d2d57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c08789031215614d6357600080fd5b863567ffffffffffffffff80821115614d7b57600080fd5b614d878a838b01614c2e565b97506020890135915080821115614d9d57600080fd5b614da98a838b01614c2e565b9650614db760408a01614ca2565b95506060890135915080821115614dcd57600080fd5b614dd98a838b01614cf9565b9450614de760808a016149c9565b935060a0890135915080821115614dfd57600080fd5b50614e0a89828a01614cf9565b9150509295509295509295565b60008060408385031215614e2a57600080fd5b8235614e35816149b3565b9150602083013560048110614e4957600080fd5b809150509250929050565b60008083601f840112614e6657600080fd5b50813567ffffffffffffffff811115614e7e57600080fd5b6020830191508360208260061b85010111156111a257600080fd5b60008060008060408587031215614eaf57600080fd5b843567ffffffffffffffff80821115614ec757600080fd5b614ed388838901614e54565b90965094506020870135915080821115614eec57600080fd5b50614ef987828801614e54565b95989497509550505050565b8015158114611d1757600080fd5b80356149d481614f05565b600082601f830112614f2f57600080fd5b81356020614f3f614c4f83614bdd565b82815260069290921b84018101918181019086841115614f5e57600080fd5b8286015b84811015614c975760408189031215614f7b5760008081fd5b614f83614b1e565b8135614f8e81614c01565b81528185013585820152835291830191604001614f62565b6000614fb4614c4f84614bdd565b8381529050602080820190600585901b840186811115614fd357600080fd5b845b8181101561500f57803567ffffffffffffffff811115614ff55760008081fd5b61500189828901614cf9565b855250928201928201614fd5565b505050509392505050565b600082601f83011261502b57600080fd5b6112c083833560208501614fa6565b60006101a0828403121561504d57600080fd5b615055614b47565b9050615060826149c9565b815261506e60208301614c23565b602082015261507f60408301614c23565b6040820152615090606083016149c9565b6060820152608082013560808201526150ab60a08301614f13565b60a08201526150bc60c083016149c9565b60c08201526150cd60e08301614c23565b60e082015261010082810135908201526101208083013567ffffffffffffffff808211156150fa57600080fd5b61510686838701614cf9565b8385015261014092508285013591508082111561512257600080fd5b61512e86838701614f1e565b8385015261016092508285013591508082111561514a57600080fd5b506151578582860161501a565b82840152505061018080830135818301525092915050565b6000806040838503121561518257600080fd5b823567ffffffffffffffff8082111561519a57600080fd5b6151a68683870161503a565b935060208501359150808211156151bc57600080fd5b506151c98582860161501a565b9150509250929050565b6151dd8184614a25565b604060208201526000611c586040830184614a92565b60006020828403121561520557600080fd5b81356112c081614c01565b600081518084526020808501945080840160005b8381101561525657815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101615224565b509495945050505050565b6020815260006112c06020830184615210565b6020808252825182820181905260009190848201906040850190845b818110156152c257835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101615290565b50909695505050505050565b60008083601f8401126152e057600080fd5b50813567ffffffffffffffff8111156152f857600080fd5b6020830191508360208285010111156111a257600080fd5b60008083601f84011261532257600080fd5b50813567ffffffffffffffff81111561533a57600080fd5b6020830191508360208260051b85010111156111a257600080fd5b60008060008060008060008060a0898b03121561537157600080fd5b883567ffffffffffffffff8082111561538957600080fd5b6153958c838d01614f1e565b995060208b01359150808211156153ab57600080fd5b6153b78c838d016152ce565b909950975060408b013591506153cc82614c01565b90955060608a013590808211156153e257600080fd5b6153ee8c838d01615310565b909650945060808b013591508082111561540757600080fd5b506154148b828c01615310565b999c989b5096995094979396929594505050565b600081518084526020808501945080840160005b83811015615256578151805173ffffffffffffffffffffffffffffffffffffffff168852830151838801526040909601959082019060010161543c565b6020815260006112c06020830184615428565b600082601f83011261549d57600080fd5b813560206154ad614c4f83614bdd565b82815260059290921b840181019181810190868411156154cc57600080fd5b8286015b84811015614c9757803567ffffffffffffffff8111156154f05760008081fd5b6154fe8986838b010161501a565b8452509183019183016154d0565b600082601f83011261551d57600080fd5b8135602061552d614c4f83614bdd565b82815260059290921b8401810191818101908684111561554c57600080fd5b8286015b84811015614c975780358352918301918301615550565b60006080828403121561557957600080fd5b615581614b6b565b9050813567ffffffffffffffff8082111561559b57600080fd5b818401915084601f8301126155af57600080fd5b813560206155bf614c4f83614bdd565b82815260059290921b840181019181810190888411156155de57600080fd5b8286015b84811015615616578035868111156155fa5760008081fd5b6156088b86838b010161503a565b8452509183019183016155e2565b508652508581013593508284111561562d57600080fd5b6156398785880161548c565b9085015250604084013591508082111561565257600080fd5b5061565f8482850161550c565b6040830152506060820135606082015292915050565b6000806040838503121561568857600080fd5b823567ffffffffffffffff808211156156a057600080fd5b6156ac86838701615567565b93506020915081850135818111156156c357600080fd5b85019050601f810186136156d657600080fd5b80356156e4614c4f82614bdd565b81815260059190911b8201830190838101908883111561570357600080fd5b928401925b8284101561572157833582529284019290840190615708565b80955050505050509250929050565b60a081016108ac828463ffffffff808251168352602082015173ffffffffffffffffffffffffffffffffffffffff8082166020860152806040850151166040860152505061ffff6060830151166060840152806080830151166080840152505050565b6000602082840312156157a557600080fd5b813567ffffffffffffffff8111156157bc57600080fd5b820160a081850312156112c057600080fd5b60008060008060008060008060e0898b0312156157ea57600080fd5b606089018a8111156157fb57600080fd5b8998503567ffffffffffffffff8082111561581557600080fd5b6158218c838d016152ce565b909950975060808b013591508082111561583a57600080fd5b6158468c838d01615310565b909750955060a08b013591508082111561585f57600080fd5b5061586c8b828c01615310565b999c989b50969995989497949560c00135949350505050565b6000806020838503121561589857600080fd5b823567ffffffffffffffff8111156158af57600080fd5b6158bb858286016152ce565b90969095509350505050565b80356fffffffffffffffffffffffffffffffff811681146149d457600080fd5b6000606082840312156158f957600080fd5b6040516060810181811067ffffffffffffffff8211171561591c5761591c614aef565b604052823561592a81614f05565b8152615938602084016158c7565b6020820152615949604084016158c7565b60408201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156108ac576108ac615955565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff808416806159e1576159e1615997565b92169190910692915050565b80820281158282048414176108ac576108ac615955565b600067ffffffffffffffff80841680615a1f57615a1f615997565b92169190910492915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615a8b57615a8b615955565b5060010190565b600063ffffffff808316818103615aab57615aab615955565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152615ae58184018a615210565b90508281036080840152615af98189615210565b905060ff871660a084015282810360c0840152615b168187614a92565b905067ffffffffffffffff851660e0840152828103610100840152615b3b8185614a92565b9c9b505050505050505050505050565b600060208284031215615b5d57600080fd5b81516112c081614c01565b60006112c0368484614fa6565b600060208284031215615b8757600080fd5b81516112c0816149b3565b808201808211156108ac576108ac615955565b608081528451608082015267ffffffffffffffff60208601511660a08201526000604086015160a060c0840152615be0610120840182614a92565b905060608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160e0860152615c1c8383614a92565b925060808901519150808584030161010086015250615c3b8282615428565b92505050615c4f602083018661ffff169052565b836040830152612cd0606083018473ffffffffffffffffffffffffffffffffffffffff169052565b60008060408385031215615c8a57600080fd5b8251615c9581614f05565b602084015190925067ffffffffffffffff811115615cb257600080fd5b8301601f81018513615cc357600080fd5b8051615cd1614c4f82614cb3565b818152866020838501011115615ce657600080fd5b615cf7826020830160208601614a6e565b8093505050509250929050565b805163ffffffff811681146149d457600080fd5b600060a08284031215615d2a57600080fd5b60405160a0810181811067ffffffffffffffff82111715615d4d57615d4d614aef565b604052615d5983615d04565b81526020830151615d6981614c01565b60208201526040830151615d7c81614c01565b6040820152606083015161ffff81168114615d9657600080fd5b6060820152615da760808401615d04565b60808201529392505050565b6101608101615e25828573ffffffffffffffffffffffffffffffffffffffff808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015250508060608301511660608401528060808301511660808401528060a08301511660a0840152505050565b825163ffffffff90811660c0840152602084015173ffffffffffffffffffffffffffffffffffffffff90811660e0850152604085015116610100840152606084015161ffff166101208401526080840151166101408301526112c0565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152615ec98285018b615210565b91508382036080850152615edd828a615210565b915060ff881660a085015283820360c0850152615efa8288614a92565b90861660e08501528381036101008501529050615b3b8185614a92565b600081518084526020808501808196508360051b8101915082860160005b85811015615f5f578284038952615f4d848351614a92565b98850198935090840190600101615f35565b5091979650505050505050565b60408152615f8760408201845167ffffffffffffffff169052565b60006020840151615fb0606084018273ffffffffffffffffffffffffffffffffffffffff169052565b50604084015173ffffffffffffffffffffffffffffffffffffffff8116608084015250606084015167ffffffffffffffff811660a084015250608084015160c083015260a084015180151560e08401525060c084015161010061601e8185018367ffffffffffffffff169052565b60e086015191506101206160498186018473ffffffffffffffffffffffffffffffffffffffff169052565b81870151925061014091508282860152808701519250506101a061016081818701526160796101e0870185614a92565b93508288015192507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06101808188870301818901526160b88686615428565b9550828a015194508188870301848901526160d38686615f17565b9550808a01516101c089015250505050508281036020840152612cd08185615f17565b6000815160208301517fffffffff000000000000000000000000000000000000000000000000000000008082169350600483101561613e5780818460040360031b1b83161693505b505050919050565b6040815260006161596040830185614a92565b8281036020840152612cd08185614a92565b60a08152600061617e60a0830188614a92565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015267ffffffffffffffff8516606084015282810360808401526161c28185614a92565b98975050505050505050565b6000602082840312156161e057600080fd5b81516112c081614f05565b600081518084526020808501945080840160005b83811015615256578151875295820195908201906001016161ff565b60608152600061622e60608301866161eb565b828103602084015261624081866161eb565b915050826040830152949350505050565b60006020828403121561626357600080fd5b5051919050565b67ffffffffffffffff81811683821601908082111561482c5761482c615955565b67ffffffffffffffff83168152604081016112c06020830184614a25565b600067ffffffffffffffff808316818103615aab57615aab615955565b6000602082840312156162d857600080fd5b813567ffffffffffffffff8111156162ef57600080fd5b611c5884828501615567565b60006040828403121561630d57600080fd5b616315614b1e565b825177ffffffffffffffffffffffffffffffffffffffffffffffff8116811461633d57600080fd5b8152602083015161634d816149b3565b60208201529392505050565b6020815260006112c06020830184615f17565b60008261637b5761637b615997565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var EVM2EVMOffRampHelperABI = EVM2EVMOffRampHelperMetaData.ABI

var EVM2EVMOffRampHelperBin = EVM2EVMOffRampHelperMetaData.Bin

func DeployEVM2EVMOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMOffRampStaticConfig, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig RateLimiterConfig) (common.Address, *types.Transaction, *EVM2EVMOffRampHelper, error) {
	parsed, err := EVM2EVMOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOffRampHelperBin), backend, staticConfig, sourceTokens, pools, rateLimiterConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMOffRampHelper{EVM2EVMOffRampHelperCaller: EVM2EVMOffRampHelperCaller{contract: contract}, EVM2EVMOffRampHelperTransactor: EVM2EVMOffRampHelperTransactor{contract: contract}, EVM2EVMOffRampHelperFilterer: EVM2EVMOffRampHelperFilterer{contract: contract}}, nil
}

type EVM2EVMOffRampHelper struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMOffRampHelperCaller
	EVM2EVMOffRampHelperTransactor
	EVM2EVMOffRampHelperFilterer
}

type EVM2EVMOffRampHelperCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampHelperTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampHelperFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampHelperSession struct {
	Contract     *EVM2EVMOffRampHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMOffRampHelperCallerSession struct {
	Contract *EVM2EVMOffRampHelperCaller
	CallOpts bind.CallOpts
}

type EVM2EVMOffRampHelperTransactorSession struct {
	Contract     *EVM2EVMOffRampHelperTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMOffRampHelperRaw struct {
	Contract *EVM2EVMOffRampHelper
}

type EVM2EVMOffRampHelperCallerRaw struct {
	Contract *EVM2EVMOffRampHelperCaller
}

type EVM2EVMOffRampHelperTransactorRaw struct {
	Contract *EVM2EVMOffRampHelperTransactor
}

func NewEVM2EVMOffRampHelper(address common.Address, backend bind.ContractBackend) (*EVM2EVMOffRampHelper, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMOffRampHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelper{address: address, abi: abi, EVM2EVMOffRampHelperCaller: EVM2EVMOffRampHelperCaller{contract: contract}, EVM2EVMOffRampHelperTransactor: EVM2EVMOffRampHelperTransactor{contract: contract}, EVM2EVMOffRampHelperFilterer: EVM2EVMOffRampHelperFilterer{contract: contract}}, nil
}

func NewEVM2EVMOffRampHelperCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMOffRampHelperCaller, error) {
	contract, err := bindEVM2EVMOffRampHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperCaller{contract: contract}, nil
}

func NewEVM2EVMOffRampHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMOffRampHelperTransactor, error) {
	contract, err := bindEVM2EVMOffRampHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperTransactor{contract: contract}, nil
}

func NewEVM2EVMOffRampHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMOffRampHelperFilterer, error) {
	contract, err := bindEVM2EVMOffRampHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperFilterer{contract: contract}, nil
}

func bindEVM2EVMOffRampHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVM2EVMOffRampHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMOffRampHelper.Contract.EVM2EVMOffRampHelperCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.EVM2EVMOffRampHelperTransactor.contract.Transfer(opts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.EVM2EVMOffRampHelperTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMOffRampHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.contract.Transfer(opts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) ISSCRIPT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "IS_SCRIPT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ISSCRIPT() (bool, error) {
	return _EVM2EVMOffRampHelper.Contract.ISSCRIPT(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) ISSCRIPT() (bool, error) {
	return _EVM2EVMOffRampHelper.Contract.ISSCRIPT(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _EVM2EVMOffRampHelper.Contract.CcipReceive(&_EVM2EVMOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _EVM2EVMOffRampHelper.Contract.CcipReceive(&_EVM2EVMOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMOffRampHelper.Contract.CurrentRateLimiterState(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMOffRampHelper.Contract.CurrentRateLimiterState(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDestinationToken(&_EVM2EVMOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDestinationToken(&_EVM2EVMOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDestinationTokens(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDestinationTokens(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOffRampDynamicConfig, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(EVM2EVMOffRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOffRampDynamicConfig)).(*EVM2EVMOffRampDynamicConfig)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetDynamicConfig() (EVM2EVMOffRampDynamicConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDynamicConfig(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetDynamicConfig() (EVM2EVMOffRampDynamicConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDynamicConfig(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMOffRampHelper.Contract.GetExecutionState(&_EVM2EVMOffRampHelper.CallOpts, sequenceNumber)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMOffRampHelper.Contract.GetExecutionState(&_EVM2EVMOffRampHelper.CallOpts, sequenceNumber)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetExecutionStateBitMap(opts *bind.CallOpts, bitmapIndex uint64) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getExecutionStateBitMap", bitmapIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetExecutionStateBitMap(bitmapIndex uint64) (*big.Int, error) {
	return _EVM2EVMOffRampHelper.Contract.GetExecutionStateBitMap(&_EVM2EVMOffRampHelper.CallOpts, bitmapIndex)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetExecutionStateBitMap(bitmapIndex uint64) (*big.Int, error) {
	return _EVM2EVMOffRampHelper.Contract.GetExecutionStateBitMap(&_EVM2EVMOffRampHelper.CallOpts, bitmapIndex)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getPoolByDestToken", destToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPoolByDestToken(&_EVM2EVMOffRampHelper.CallOpts, destToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPoolByDestToken(&_EVM2EVMOffRampHelper.CallOpts, destToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPoolBySourceToken(&_EVM2EVMOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPoolBySourceToken(&_EVM2EVMOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getSenderNonce", sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMOffRampHelper.Contract.GetSenderNonce(&_EVM2EVMOffRampHelper.CallOpts, sender)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMOffRampHelper.Contract.GetSenderNonce(&_EVM2EVMOffRampHelper.CallOpts, sender)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOffRampStaticConfig, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(EVM2EVMOffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOffRampStaticConfig)).(*EVM2EVMOffRampStaticConfig)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetStaticConfig() (EVM2EVMOffRampStaticConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetStaticConfig(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetStaticConfig() (EVM2EVMOffRampStaticConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetStaticConfig(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetSupportedTokens(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetSupportedTokens(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetTokenLimitAdmin(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetTokenLimitAdmin(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetTransmitters(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetTransmitters(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMOffRampHelper.Contract.LatestConfigDetails(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMOffRampHelper.Contract.LatestConfigDetails(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) MetadataHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "metadataHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) MetadataHash() ([32]byte, error) {
	return _EVM2EVMOffRampHelper.Contract.MetadataHash(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) MetadataHash() ([32]byte, error) {
	return _EVM2EVMOffRampHelper.Contract.MetadataHash(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Owner() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.Owner(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.Owner(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) TypeAndVersion() (string, error) {
	return _EVM2EVMOffRampHelper.Contract.TypeAndVersion(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMOffRampHelper.Contract.TypeAndVersion(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.AcceptOwnership(&_EVM2EVMOffRampHelper.TransactOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.AcceptOwnership(&_EVM2EVMOffRampHelper.TransactOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "applyPoolUpdates", removes, adds)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ApplyPoolUpdates(removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ApplyPoolUpdates(&_EVM2EVMOffRampHelper.TransactOpts, removes, adds)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ApplyPoolUpdates(removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ApplyPoolUpdates(&_EVM2EVMOffRampHelper.TransactOpts, removes, adds)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Execute(opts *bind.TransactOpts, rep InternalExecutionReport, manualExecGasLimits []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "execute", rep, manualExecGasLimits)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Execute(rep InternalExecutionReport, manualExecGasLimits []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Execute(&_EVM2EVMOffRampHelper.TransactOpts, rep, manualExecGasLimits)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Execute(rep InternalExecutionReport, manualExecGasLimits []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Execute(&_EVM2EVMOffRampHelper.TransactOpts, rep, manualExecGasLimits)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport, gasLimitOverrides []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "manuallyExecute", report, gasLimitOverrides)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ManuallyExecute(report InternalExecutionReport, gasLimitOverrides []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMOffRampHelper.TransactOpts, report, gasLimitOverrides)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ManuallyExecute(report InternalExecutionReport, gasLimitOverrides []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMOffRampHelper.TransactOpts, report, gasLimitOverrides)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ReleaseOrMintTokens(opts *bind.TransactOpts, sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, offchainTokenData [][]byte, sourceTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "releaseOrMintTokens", sourceTokenAmounts, originalSender, receiver, offchainTokenData, sourceTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ReleaseOrMintTokens(sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, offchainTokenData [][]byte, sourceTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintTokens(&_EVM2EVMOffRampHelper.TransactOpts, sourceTokenAmounts, originalSender, receiver, offchainTokenData, sourceTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ReleaseOrMintTokens(sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, offchainTokenData [][]byte, sourceTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintTokens(&_EVM2EVMOffRampHelper.TransactOpts, sourceTokenAmounts, originalSender, receiver, offchainTokenData, sourceTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "report", executableMessages)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Report(&_EVM2EVMOffRampHelper.TransactOpts, executableMessages)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Report(&_EVM2EVMOffRampHelper.TransactOpts, executableMessages)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setAdmin", newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetAdmin(&_EVM2EVMOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetAdmin(&_EVM2EVMOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetExecutionStateHelper(opts *bind.TransactOpts, sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setExecutionStateHelper", sequenceNumber, state)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetExecutionStateHelper(sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetExecutionStateHelper(&_EVM2EVMOffRampHelper.TransactOpts, sequenceNumber, state)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetExecutionStateHelper(sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetExecutionStateHelper(&_EVM2EVMOffRampHelper.TransactOpts, sequenceNumber, state)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setOCR2Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetOCR2Config(&_EVM2EVMOffRampHelper.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetOCR2Config(&_EVM2EVMOffRampHelper.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.TransferOwnership(&_EVM2EVMOffRampHelper.TransactOpts, to)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.TransferOwnership(&_EVM2EVMOffRampHelper.TransactOpts, to)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Transmit(&_EVM2EVMOffRampHelper.TransactOpts, reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Transmit(&_EVM2EVMOffRampHelper.TransactOpts, reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) TrialExecute(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "trialExecute", message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) TrialExecute(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.TrialExecute(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) TrialExecute(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.TrialExecute(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData)
}

type EVM2EVMOffRampHelperAdminSetIterator struct {
	Event *EVM2EVMOffRampHelperAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperAdminSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperAdminSet struct {
	NewAdmin common.Address
	Raw      types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperAdminSetIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperAdminSetIterator{contract: _EVM2EVMOffRampHelper.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperAdminSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperAdminSet)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "AdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseAdminSet(log types.Log) (*EVM2EVMOffRampHelperAdminSet, error) {
	event := new(EVM2EVMOffRampHelperAdminSet)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperConfigSetIterator struct {
	Event *EVM2EVMOffRampHelperConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperConfigSet struct {
	StaticConfig  EVM2EVMOffRampStaticConfig
	DynamicConfig EVM2EVMOffRampDynamicConfig
	Raw           types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperConfigSetIterator{contract: _EVM2EVMOffRampHelper.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperConfigSet)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseConfigSet(log types.Log) (*EVM2EVMOffRampHelperConfigSet, error) {
	event := new(EVM2EVMOffRampHelperConfigSet)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperConfigSet0Iterator struct {
	Event *EVM2EVMOffRampHelperConfigSet0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperConfigSet0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperConfigSet0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperConfigSet0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperConfigSet0Iterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperConfigSet0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperConfigSet0 struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterConfigSet0(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSet0Iterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "ConfigSet0")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperConfigSet0Iterator{contract: _EVM2EVMOffRampHelper.contract, event: "ConfigSet0", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchConfigSet0(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet0) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "ConfigSet0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperConfigSet0)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigSet0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseConfigSet0(log types.Log) (*EVM2EVMOffRampHelperConfigSet0, error) {
	event := new(EVM2EVMOffRampHelperConfigSet0)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigSet0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperExecutionStateChangedIterator struct {
	Event *EVM2EVMOffRampHelperExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperExecutionStateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperExecutionStateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperExecutionStateChanged struct {
	SequenceNumber uint64
	MessageId      [32]byte
	State          uint8
	ReturnData     []byte
	Raw            types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMOffRampHelperExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperExecutionStateChangedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperExecutionStateChanged)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMOffRampHelperExecutionStateChanged, error) {
	event := new(EVM2EVMOffRampHelperExecutionStateChanged)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMOffRampHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperOwnershipTransferRequestedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperOwnershipTransferRequested)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferRequested, error) {
	event := new(EVM2EVMOffRampHelperOwnershipTransferRequested)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperOwnershipTransferredIterator struct {
	Event *EVM2EVMOffRampHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperOwnershipTransferredIterator{contract: _EVM2EVMOffRampHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperOwnershipTransferred)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferred, error) {
	event := new(EVM2EVMOffRampHelperOwnershipTransferred)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperPoolAddedIterator struct {
	Event *EVM2EVMOffRampHelperPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperPoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperPoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperPoolAddedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperPoolAdded)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMOffRampHelperPoolAdded, error) {
	event := new(EVM2EVMOffRampHelperPoolAdded)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperPoolRemovedIterator struct {
	Event *EVM2EVMOffRampHelperPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperPoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperPoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperPoolRemovedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperPoolRemoved)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMOffRampHelperPoolRemoved, error) {
	event := new(EVM2EVMOffRampHelperPoolRemoved)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperSkippedIncorrectNonceIterator struct {
	Event *EVM2EVMOffRampHelperSkippedIncorrectNonce

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperSkippedIncorrectNonceIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperSkippedIncorrectNonce)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperSkippedIncorrectNonce)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperSkippedIncorrectNonceIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperSkippedIncorrectNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperSkippedIncorrectNonce struct {
	Nonce  uint64
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedIncorrectNonceIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "SkippedIncorrectNonce", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperSkippedIncorrectNonceIterator{contract: _EVM2EVMOffRampHelper.contract, event: "SkippedIncorrectNonce", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "SkippedIncorrectNonce", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperSkippedIncorrectNonce)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampHelperSkippedIncorrectNonce, error) {
	event := new(EVM2EVMOffRampHelperSkippedIncorrectNonce)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator struct {
	Event *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight struct {
	Nonce  uint64
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "SkippedSenderWithPreviousRampMessageInflight", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator{contract: _EVM2EVMOffRampHelper.contract, event: "SkippedSenderWithPreviousRampMessageInflight", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight, nonce []uint64, sender []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "SkippedSenderWithPreviousRampMessageInflight", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "SkippedSenderWithPreviousRampMessageInflight", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight, error) {
	event := new(EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "SkippedSenderWithPreviousRampMessageInflight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperTransmittedIterator struct {
	Event *EVM2EVMOffRampHelperTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperTransmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOffRampHelperTransmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOffRampHelperTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperTransmittedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperTransmitted)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseTransmitted(log types.Log) (*EVM2EVMOffRampHelperTransmitted, error) {
	event := new(EVM2EVMOffRampHelperTransmitted)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMOffRampHelper.abi.Events["AdminSet"].ID:
		return _EVM2EVMOffRampHelper.ParseAdminSet(log)
	case _EVM2EVMOffRampHelper.abi.Events["ConfigSet"].ID:
		return _EVM2EVMOffRampHelper.ParseConfigSet(log)
	case _EVM2EVMOffRampHelper.abi.Events["ConfigSet0"].ID:
		return _EVM2EVMOffRampHelper.ParseConfigSet0(log)
	case _EVM2EVMOffRampHelper.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMOffRampHelper.ParseExecutionStateChanged(log)
	case _EVM2EVMOffRampHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMOffRampHelper.ParseOwnershipTransferRequested(log)
	case _EVM2EVMOffRampHelper.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMOffRampHelper.ParseOwnershipTransferred(log)
	case _EVM2EVMOffRampHelper.abi.Events["PoolAdded"].ID:
		return _EVM2EVMOffRampHelper.ParsePoolAdded(log)
	case _EVM2EVMOffRampHelper.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMOffRampHelper.ParsePoolRemoved(log)
	case _EVM2EVMOffRampHelper.abi.Events["SkippedIncorrectNonce"].ID:
		return _EVM2EVMOffRampHelper.ParseSkippedIncorrectNonce(log)
	case _EVM2EVMOffRampHelper.abi.Events["SkippedSenderWithPreviousRampMessageInflight"].ID:
		return _EVM2EVMOffRampHelper.ParseSkippedSenderWithPreviousRampMessageInflight(log)
	case _EVM2EVMOffRampHelper.abi.Events["Transmitted"].ID:
		return _EVM2EVMOffRampHelper.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMOffRampHelperAdminSet) Topic() common.Hash {
	return common.HexToHash("0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c")
}

func (EVM2EVMOffRampHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x737ef22d3f6615e342ed21c69e06620dbc5c8a261ed7cfb2ce214806b1f76eda")
}

func (EVM2EVMOffRampHelperConfigSet0) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMOffRampHelperExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")
}

func (EVM2EVMOffRampHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMOffRampHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMOffRampHelperPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMOffRampHelperPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMOffRampHelperSkippedIncorrectNonce) Topic() common.Hash {
	return common.HexToHash("0xd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf41237")
}

func (EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight) Topic() common.Hash {
	return common.HexToHash("0xe44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5d")
}

func (EVM2EVMOffRampHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelper) Address() common.Address {
	return _EVM2EVMOffRampHelper.address
}

type EVM2EVMOffRampHelperInterface interface {
	ISSCRIPT(opts *bind.CallOpts) (bool, error)

	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetExecutionStateBitMap(opts *bind.CallOpts, bitmapIndex uint64) (*big.Int, error)

	GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOffRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	MetadataHash(opts *bind.CallOpts) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, rep InternalExecutionReport, manualExecGasLimits []*big.Int) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport, gasLimitOverrides []*big.Int) (*types.Transaction, error)

	ReleaseOrMintTokens(opts *bind.TransactOpts, sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, offchainTokenData [][]byte, sourceTokenData [][]byte) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetExecutionStateHelper(opts *bind.TransactOpts, sequenceNumber uint64, state uint8) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error)

	TrialExecute(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMOffRampHelperAdminSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMOffRampHelperConfigSet, error)

	FilterConfigSet0(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSet0Iterator, error)

	WatchConfigSet0(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet0) (event.Subscription, error)

	ParseConfigSet0(log types.Log) (*EVM2EVMOffRampHelperConfigSet0, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMOffRampHelperExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMOffRampHelperExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferred, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMOffRampHelperPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMOffRampHelperPoolRemoved, error)

	FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedIncorrectNonceIterator, error)

	WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampHelperSkippedIncorrectNonce, error)

	FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator, error)

	WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMOffRampHelperTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
