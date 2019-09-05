// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/smartcontractkit/chainlink/core/store (interfaces: EthClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	go_ethereum "github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	assets "github.com/smartcontractkit/chainlink/core/store/assets"
	models "github.com/smartcontractkit/chainlink/core/store/models"
	big "math/big"
	reflect "reflect"
)

// MockEthClient is a mock of EthClient interface
type MockEthClient struct {
	ctrl     *gomock.Controller
	recorder *MockEthClientMockRecorder
}

// MockEthClientMockRecorder is the mock recorder for MockEthClient
type MockEthClientMockRecorder struct {
	mock *MockEthClient
}

// NewMockEthClient creates a new mock instance
func NewMockEthClient(ctrl *gomock.Controller) *MockEthClient {
	mock := &MockEthClient{ctrl: ctrl}
	mock.recorder = &MockEthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEthClient) EXPECT() *MockEthClientMockRecorder {
	return m.recorder
}

// GetBlockByNumber mocks base method
func (m *MockEthClient) GetBlockByNumber(arg0 string) (models.BlockHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0)
	ret0, _ := ret[0].(models.BlockHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber
func (mr *MockEthClientMockRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockEthClient)(nil).GetBlockByNumber), arg0)
}

// GetChainID mocks base method
func (m *MockEthClient) GetChainID() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainID")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainID indicates an expected call of GetChainID
func (mr *MockEthClientMockRecorder) GetChainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainID", reflect.TypeOf((*MockEthClient)(nil).GetChainID))
}

// GetERC20Balance mocks base method
func (m *MockEthClient) GetERC20Balance(arg0, arg1 common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetERC20Balance", arg0, arg1)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetERC20Balance indicates an expected call of GetERC20Balance
func (mr *MockEthClientMockRecorder) GetERC20Balance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetERC20Balance", reflect.TypeOf((*MockEthClient)(nil).GetERC20Balance), arg0, arg1)
}

// GetEthBalance mocks base method
func (m *MockEthClient) GetEthBalance(arg0 common.Address) (*assets.Eth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEthBalance", arg0)
	ret0, _ := ret[0].(*assets.Eth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEthBalance indicates an expected call of GetEthBalance
func (mr *MockEthClientMockRecorder) GetEthBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEthBalance", reflect.TypeOf((*MockEthClient)(nil).GetEthBalance), arg0)
}

// GetLogs mocks base method
func (m *MockEthClient) GetLogs(arg0 go_ethereum.FilterQuery) ([]models.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0)
	ret0, _ := ret[0].([]models.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockEthClientMockRecorder) GetLogs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockEthClient)(nil).GetLogs), arg0)
}

// GetNonce mocks base method
func (m *MockEthClient) GetNonce(arg0 common.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNonce indicates an expected call of GetNonce
func (mr *MockEthClientMockRecorder) GetNonce(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockEthClient)(nil).GetNonce), arg0)
}

// GetTxReceipt mocks base method
func (m *MockEthClient) GetTxReceipt(arg0 common.Hash) (*models.TxReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxReceipt", arg0)
	ret0, _ := ret[0].(*models.TxReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxReceipt indicates an expected call of GetTxReceipt
func (mr *MockEthClientMockRecorder) GetTxReceipt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxReceipt", reflect.TypeOf((*MockEthClient)(nil).GetTxReceipt), arg0)
}

// SendRawTx mocks base method
func (m *MockEthClient) SendRawTx(arg0 string) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRawTx", arg0)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTx indicates an expected call of SendRawTx
func (mr *MockEthClientMockRecorder) SendRawTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTx", reflect.TypeOf((*MockEthClient)(nil).SendRawTx), arg0)
}

// SubscribeToLogs mocks base method
func (m *MockEthClient) SubscribeToLogs(arg0 chan<- models.Log, arg1 go_ethereum.FilterQuery) (models.EthSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToLogs", arg0, arg1)
	ret0, _ := ret[0].(models.EthSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToLogs indicates an expected call of SubscribeToLogs
func (mr *MockEthClientMockRecorder) SubscribeToLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToLogs", reflect.TypeOf((*MockEthClient)(nil).SubscribeToLogs), arg0, arg1)
}

// SubscribeToNewHeads mocks base method
func (m *MockEthClient) SubscribeToNewHeads(arg0 chan<- models.BlockHeader) (models.EthSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToNewHeads", arg0)
	ret0, _ := ret[0].(models.EthSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToNewHeads indicates an expected call of SubscribeToNewHeads
func (mr *MockEthClientMockRecorder) SubscribeToNewHeads(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNewHeads", reflect.TypeOf((*MockEthClient)(nil).SubscribeToNewHeads), arg0)
}
