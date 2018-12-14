// Code generated by MockGen. DO NOT EDIT.
// Source: store/tx_manager.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	go_ethereum "github.com/ethereum/go-ethereum"
	accounts "github.com/ethereum/go-ethereum/accounts"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	store "github.com/smartcontractkit/chainlink/store"
	assets "github.com/smartcontractkit/chainlink/store/assets"
	models "github.com/smartcontractkit/chainlink/store/models"
	big "math/big"
	reflect "reflect"
)

// MockTxManager is a mock of TxManager interface
type MockTxManager struct {
	ctrl     *gomock.Controller
	recorder *MockTxManagerMockRecorder
}

// MockTxManagerMockRecorder is the mock recorder for MockTxManager
type MockTxManagerMockRecorder struct {
	mock *MockTxManager
}

// NewMockTxManager creates a new mock instance
func NewMockTxManager(ctrl *gomock.Controller) *MockTxManager {
	mock := &MockTxManager{ctrl: ctrl}
	mock.recorder = &MockTxManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxManager) EXPECT() *MockTxManagerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockTxManager) Start(accounts []accounts.Account) error {
	ret := m.ctrl.Call(m, "Start", accounts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockTxManagerMockRecorder) Start(accounts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTxManager)(nil).Start), accounts)
}

// CreateTx mocks base method
func (m *MockTxManager) CreateTx(to common.Address, data []byte) (*models.Tx, error) {
	ret := m.ctrl.Call(m, "CreateTx", to, data)
	ret0, _ := ret[0].(*models.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx
func (mr *MockTxManagerMockRecorder) CreateTx(to, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockTxManager)(nil).CreateTx), to, data)
}

// CreateTxWithGas mocks base method
func (m *MockTxManager) CreateTxWithGas(to common.Address, data []byte, gasPriceWei *big.Int, gasLimit uint64) (*models.Tx, error) {
	ret := m.ctrl.Call(m, "CreateTxWithGas", to, data, gasPriceWei, gasLimit)
	ret0, _ := ret[0].(*models.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTxWithGas indicates an expected call of CreateTxWithGas
func (mr *MockTxManagerMockRecorder) CreateTxWithGas(to, data, gasPriceWei, gasLimit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTxWithGas", reflect.TypeOf((*MockTxManager)(nil).CreateTxWithGas), to, data, gasPriceWei, gasLimit)
}

// MeetsMinConfirmations mocks base method
func (m *MockTxManager) MeetsMinConfirmations(hash common.Hash) (bool, error) {
	ret := m.ctrl.Call(m, "MeetsMinConfirmations", hash)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MeetsMinConfirmations indicates an expected call of MeetsMinConfirmations
func (mr *MockTxManagerMockRecorder) MeetsMinConfirmations(hash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeetsMinConfirmations", reflect.TypeOf((*MockTxManager)(nil).MeetsMinConfirmations), hash)
}

// ContractLINKBalance mocks base method
func (m *MockTxManager) ContractLINKBalance(wr models.WithdrawalRequest) (assets.Link, error) {
	ret := m.ctrl.Call(m, "ContractLINKBalance", wr)
	ret0, _ := ret[0].(assets.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContractLINKBalance indicates an expected call of ContractLINKBalance
func (mr *MockTxManagerMockRecorder) ContractLINKBalance(wr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContractLINKBalance", reflect.TypeOf((*MockTxManager)(nil).ContractLINKBalance), wr)
}

// WithdrawLINK mocks base method
func (m *MockTxManager) WithdrawLINK(wr models.WithdrawalRequest) (common.Hash, error) {
	ret := m.ctrl.Call(m, "WithdrawLINK", wr)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithdrawLINK indicates an expected call of WithdrawLINK
func (mr *MockTxManagerMockRecorder) WithdrawLINK(wr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawLINK", reflect.TypeOf((*MockTxManager)(nil).WithdrawLINK), wr)
}

// GetLINKBalance mocks base method
func (m *MockTxManager) GetLINKBalance(address common.Address) (*assets.Link, error) {
	ret := m.ctrl.Call(m, "GetLINKBalance", address)
	ret0, _ := ret[0].(*assets.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLINKBalance indicates an expected call of GetLINKBalance
func (mr *MockTxManagerMockRecorder) GetLINKBalance(address interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLINKBalance", reflect.TypeOf((*MockTxManager)(nil).GetLINKBalance), address)
}

// NextActiveAccount mocks base method
func (m *MockTxManager) NextActiveAccount() *store.ManagedAccount {
	ret := m.ctrl.Call(m, "NextActiveAccount")
	ret0, _ := ret[0].(*store.ManagedAccount)
	return ret0
}

// NextActiveAccount indicates an expected call of NextActiveAccount
func (mr *MockTxManagerMockRecorder) NextActiveAccount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextActiveAccount", reflect.TypeOf((*MockTxManager)(nil).NextActiveAccount))
}

// GetEthBalance mocks base method
func (m *MockTxManager) GetEthBalance(address common.Address) (*assets.Eth, error) {
	ret := m.ctrl.Call(m, "GetEthBalance", address)
	ret0, _ := ret[0].(*assets.Eth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEthBalance indicates an expected call of GetEthBalance
func (mr *MockTxManagerMockRecorder) GetEthBalance(address interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEthBalance", reflect.TypeOf((*MockTxManager)(nil).GetEthBalance), address)
}

// SubscribeToNewHeads mocks base method
func (m *MockTxManager) SubscribeToNewHeads(channel chan<- models.BlockHeader) (models.EthSubscription, error) {
	ret := m.ctrl.Call(m, "SubscribeToNewHeads", channel)
	ret0, _ := ret[0].(models.EthSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToNewHeads indicates an expected call of SubscribeToNewHeads
func (mr *MockTxManagerMockRecorder) SubscribeToNewHeads(channel interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNewHeads", reflect.TypeOf((*MockTxManager)(nil).SubscribeToNewHeads), channel)
}

// GetBlockByNumber mocks base method
func (m *MockTxManager) GetBlockByNumber(hex string) (models.BlockHeader, error) {
	ret := m.ctrl.Call(m, "GetBlockByNumber", hex)
	ret0, _ := ret[0].(models.BlockHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber
func (mr *MockTxManagerMockRecorder) GetBlockByNumber(hex interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockTxManager)(nil).GetBlockByNumber), hex)
}

// SubscribeToLogs mocks base method
func (m *MockTxManager) SubscribeToLogs(channel chan<- store.Log, q go_ethereum.FilterQuery) (models.EthSubscription, error) {
	ret := m.ctrl.Call(m, "SubscribeToLogs", channel, q)
	ret0, _ := ret[0].(models.EthSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToLogs indicates an expected call of SubscribeToLogs
func (mr *MockTxManagerMockRecorder) SubscribeToLogs(channel, q interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToLogs", reflect.TypeOf((*MockTxManager)(nil).SubscribeToLogs), channel, q)
}

// GetLogs mocks base method
func (m *MockTxManager) GetLogs(q go_ethereum.FilterQuery) ([]store.Log, error) {
	ret := m.ctrl.Call(m, "GetLogs", q)
	ret0, _ := ret[0].([]store.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockTxManagerMockRecorder) GetLogs(q interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockTxManager)(nil).GetLogs), q)
}
