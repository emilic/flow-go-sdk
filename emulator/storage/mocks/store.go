// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dapperlabs/flow-go-sdk/emulator/storage (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	flow_go_sdk "github.com/dapperlabs/flow-go-sdk"
	types "github.com/dapperlabs/flow-go-sdk/emulator/types"
	crypto "github.com/dapperlabs/flow-go/crypto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// BlockByHash mocks base method
func (m *MockStore) BlockByHash(arg0 crypto.Hash) (types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", arg0)
	ret0, _ := ret[0].(types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash
func (mr *MockStoreMockRecorder) BlockByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockStore)(nil).BlockByHash), arg0)
}

// BlockByNumber mocks base method
func (m *MockStore) BlockByNumber(arg0 uint64) (types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", arg0)
	ret0, _ := ret[0].(types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber
func (mr *MockStoreMockRecorder) BlockByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockStore)(nil).BlockByNumber), arg0)
}

// CommitBlock mocks base method
func (m *MockStore) CommitBlock(arg0 types.Block, arg1 []flow_go_sdk.Transaction, arg2 types.LedgerDelta, arg3 []flow_go_sdk.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitBlock", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitBlock indicates an expected call of CommitBlock
func (mr *MockStoreMockRecorder) CommitBlock(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitBlock", reflect.TypeOf((*MockStore)(nil).CommitBlock), arg0, arg1, arg2, arg3)
}

// InsertBlock mocks base method
func (m *MockStore) InsertBlock(arg0 types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBlock indicates an expected call of InsertBlock
func (mr *MockStoreMockRecorder) InsertBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBlock", reflect.TypeOf((*MockStore)(nil).InsertBlock), arg0)
}

// InsertEvents mocks base method
func (m *MockStore) InsertEvents(arg0 uint64, arg1 []flow_go_sdk.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertEvents", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertEvents indicates an expected call of InsertEvents
func (mr *MockStoreMockRecorder) InsertEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEvents", reflect.TypeOf((*MockStore)(nil).InsertEvents), arg0, arg1)
}

// InsertLedgerDelta mocks base method
func (m *MockStore) InsertLedgerDelta(arg0 uint64, arg1 types.LedgerDelta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertLedgerDelta", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertLedgerDelta indicates an expected call of InsertLedgerDelta
func (mr *MockStoreMockRecorder) InsertLedgerDelta(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertLedgerDelta", reflect.TypeOf((*MockStore)(nil).InsertLedgerDelta), arg0, arg1)
}

// InsertTransaction mocks base method
func (m *MockStore) InsertTransaction(arg0 flow_go_sdk.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTransaction indicates an expected call of InsertTransaction
func (mr *MockStoreMockRecorder) InsertTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTransaction", reflect.TypeOf((*MockStore)(nil).InsertTransaction), arg0)
}

// LatestBlock mocks base method
func (m *MockStore) LatestBlock() (types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBlock")
	ret0, _ := ret[0].(types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBlock indicates an expected call of LatestBlock
func (mr *MockStoreMockRecorder) LatestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBlock", reflect.TypeOf((*MockStore)(nil).LatestBlock))
}

// LedgerViewByNumber mocks base method
func (m *MockStore) LedgerViewByNumber(arg0 uint64) *types.LedgerView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LedgerViewByNumber", arg0)
	ret0, _ := ret[0].(*types.LedgerView)
	return ret0
}

// LedgerViewByNumber indicates an expected call of LedgerViewByNumber
func (mr *MockStoreMockRecorder) LedgerViewByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LedgerViewByNumber", reflect.TypeOf((*MockStore)(nil).LedgerViewByNumber), arg0)
}

// RetrieveEvents mocks base method
func (m *MockStore) RetrieveEvents(arg0 string, arg1, arg2 uint64) ([]flow_go_sdk.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].([]flow_go_sdk.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveEvents indicates an expected call of RetrieveEvents
func (mr *MockStoreMockRecorder) RetrieveEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveEvents", reflect.TypeOf((*MockStore)(nil).RetrieveEvents), arg0, arg1, arg2)
}

// TransactionByHash mocks base method
func (m *MockStore) TransactionByHash(arg0 crypto.Hash) (flow_go_sdk.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", arg0)
	ret0, _ := ret[0].(flow_go_sdk.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByHash indicates an expected call of TransactionByHash
func (mr *MockStoreMockRecorder) TransactionByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockStore)(nil).TransactionByHash), arg0)
}
