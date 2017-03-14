// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cosminrentea/gobbler/server/auth (interfaces: AccessManager)

package auth

import (
	protocol "github.com/cosminrentea/gobbler/protocol"
	
	gomock "github.com/golang/mock/gomock"
)

// Mock of AccessManager interface
type MockAccessManager struct {
	ctrl     *gomock.Controller
	recorder *_MockAccessManagerRecorder
}

// Recorder for MockAccessManager (not exported)
type _MockAccessManagerRecorder struct {
	mock *MockAccessManager
}

func NewMockAccessManager(ctrl *gomock.Controller) *MockAccessManager {
	mock := &MockAccessManager{ctrl: ctrl}
	mock.recorder = &_MockAccessManagerRecorder{mock}
	return mock
}

func (_m *MockAccessManager) EXPECT() *_MockAccessManagerRecorder {
	return _m.recorder
}

func (_m *MockAccessManager) IsAllowed(_param0 AccessType, _param1 string, _param2 protocol.Path) bool {
	ret := _m.ctrl.Call(_m, "IsAllowed", _param0, _param1, _param2)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockAccessManagerRecorder) IsAllowed(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsAllowed", arg0, arg1, arg2)
}
