// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cosminrentea/gobbler/server/router (interfaces: Router)

package websocket

import (
	protocol "github.com/cosminrentea/gobbler/protocol"
	cluster "github.com/cosminrentea/gobbler/server/cluster"
	kvstore "github.com/cosminrentea/gobbler/server/kvstore"
	router "github.com/cosminrentea/gobbler/server/router"
	store "github.com/cosminrentea/gobbler/server/store"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *_MockRouterRecorder
}

// Recorder for MockRouter (not exported)
type _MockRouterRecorder struct {
	mock *MockRouter
}

func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &_MockRouterRecorder{mock}
	return mock
}

func (_m *MockRouter) EXPECT() *_MockRouterRecorder {
	return _m.recorder
}

func (_m *MockRouter) Cluster() *cluster.Cluster {
	ret := _m.ctrl.Call(_m, "Cluster")
	ret0, _ := ret[0].(*cluster.Cluster)
	return ret0
}

func (_mr *_MockRouterRecorder) Cluster() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Cluster")
}

func (_m *MockRouter) Done() <-chan bool {
	ret := _m.ctrl.Call(_m, "Done")
	ret0, _ := ret[0].(<-chan bool)
	return ret0
}

func (_mr *_MockRouterRecorder) Done() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Done")
}

func (_m *MockRouter) Fetch(_param0 *store.FetchRequest) error {
	ret := _m.ctrl.Call(_m, "Fetch", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRouterRecorder) Fetch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fetch", arg0)
}

func (_m *MockRouter) GetSubscribers(_param0 string) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "GetSubscribers", _param0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouterRecorder) GetSubscribers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscribers", arg0)
}

func (_m *MockRouter) HandleMessage(_param0 *protocol.Message) error {
	ret := _m.ctrl.Call(_m, "HandleMessage", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRouterRecorder) HandleMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandleMessage", arg0)
}

func (_m *MockRouter) KVStore() (kvstore.KVStore, error) {
	ret := _m.ctrl.Call(_m, "KVStore")
	ret0, _ := ret[0].(kvstore.KVStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouterRecorder) KVStore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KVStore")
}

func (_m *MockRouter) MessageStore() (store.MessageStore, error) {
	ret := _m.ctrl.Call(_m, "MessageStore")
	ret0, _ := ret[0].(store.MessageStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouterRecorder) MessageStore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MessageStore")
}

func (_m *MockRouter) Subscribe(_param0 *router.Route) (*router.Route, error) {
	ret := _m.ctrl.Call(_m, "Subscribe", _param0)
	ret0, _ := ret[0].(*router.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouterRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Subscribe", arg0)
}

func (_m *MockRouter) Unsubscribe(_param0 *router.Route) {
	_m.ctrl.Call(_m, "Unsubscribe", _param0)
}

func (_mr *_MockRouterRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unsubscribe", arg0)
}
