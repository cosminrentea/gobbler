// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cosminrentea/gobbler/server/store (interfaces: MessageStore)

package router

import (
	protocol "github.com/cosminrentea/gobbler/protocol"
	store "github.com/cosminrentea/gobbler/server/store"
	gomock "github.com/golang/mock/gomock"
)

// Mock of MessageStore interface
type MockMessageStore struct {
	ctrl     *gomock.Controller
	recorder *_MockMessageStoreRecorder
}

// Recorder for MockMessageStore (not exported)
type _MockMessageStoreRecorder struct {
	mock *MockMessageStore
}

func NewMockMessageStore(ctrl *gomock.Controller) *MockMessageStore {
	mock := &MockMessageStore{ctrl: ctrl}
	mock.recorder = &_MockMessageStoreRecorder{mock}
	return mock
}

func (_m *MockMessageStore) EXPECT() *_MockMessageStoreRecorder {
	return _m.recorder
}

func (_m *MockMessageStore) DoInTx(_param0 string, _param1 func(uint64) error) error {
	ret := _m.ctrl.Call(_m, "DoInTx", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMessageStoreRecorder) DoInTx(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoInTx", arg0, arg1)
}

func (_m *MockMessageStore) Fetch(_param0 *store.FetchRequest) {
	_m.ctrl.Call(_m, "Fetch", _param0)
}

func (_mr *_MockMessageStoreRecorder) Fetch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fetch", arg0)
}

func (_m *MockMessageStore) GenerateNextMsgID(_param0 string, _param1 byte) (uint64, int64, error) {
	ret := _m.ctrl.Call(_m, "GenerateNextMsgID", _param0, _param1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockMessageStoreRecorder) GenerateNextMsgID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateNextMsgID", arg0, arg1)
}

func (_m *MockMessageStore) MaxMessageID(_param0 string) (uint64, error) {
	ret := _m.ctrl.Call(_m, "MaxMessageID", _param0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMessageStoreRecorder) MaxMessageID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MaxMessageID", arg0)
}

func (_m *MockMessageStore) Partition(_param0 string) (store.MessagePartition, error) {
	ret := _m.ctrl.Call(_m, "Partition", _param0)
	ret0, _ := ret[0].(store.MessagePartition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMessageStoreRecorder) Partition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Partition", arg0)
}

func (_m *MockMessageStore) Partitions() ([]store.MessagePartition, error) {
	ret := _m.ctrl.Call(_m, "Partitions")
	ret0, _ := ret[0].([]store.MessagePartition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMessageStoreRecorder) Partitions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Partitions")
}

func (_m *MockMessageStore) Store(_param0 string, _param1 uint64, _param2 []byte) error {
	ret := _m.ctrl.Call(_m, "Store", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMessageStoreRecorder) Store(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Store", arg0, arg1, arg2)
}

func (_m *MockMessageStore) StoreMessage(_param0 *protocol.Message, _param1 byte) (int, error) {
	ret := _m.ctrl.Call(_m, "StoreMessage", _param0, _param1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMessageStoreRecorder) StoreMessage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StoreMessage", arg0, arg1)
}
