// Code generated by MockGen. DO NOT EDIT.
// Source: basic/basic.go

// Package mock_basic is a generated GoMock package.
package mock_basic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDb is a mock of Db interface.
type MockDb struct {
	ctrl     *gomock.Controller
	recorder *MockDbMockRecorder
}

// MockDbMockRecorder is the mock recorder for MockDb.
type MockDbMockRecorder struct {
	mock *MockDb
}

// NewMockDb creates a new mock instance.
func NewMockDb(ctrl *gomock.Controller) *MockDb {
	mock := &MockDb{ctrl: ctrl}
	mock.recorder = &MockDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDb) EXPECT() *MockDbMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDb) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDbMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDb)(nil).Get), key)
}

// Put mocks base method.
func (m *MockDb) Put(key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockDbMockRecorder) Put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockDb)(nil).Put), key, value)
}
