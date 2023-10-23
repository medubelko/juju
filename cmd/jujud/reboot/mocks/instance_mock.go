// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/environs/instances (interfaces: Instance)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	instance "github.com/juju/juju/core/instance"
	network "github.com/juju/juju/core/network"
	envcontext "github.com/juju/juju/environs/envcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockInstance is a mock of Instance interface.
type MockInstance struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMockRecorder
}

// MockInstanceMockRecorder is the mock recorder for MockInstance.
type MockInstanceMockRecorder struct {
	mock *MockInstance
}

// NewMockInstance creates a new mock instance.
func NewMockInstance(ctrl *gomock.Controller) *MockInstance {
	mock := &MockInstance{ctrl: ctrl}
	mock.recorder = &MockInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstance) EXPECT() *MockInstanceMockRecorder {
	return m.recorder
}

// Addresses mocks base method.
func (m *MockInstance) Addresses(arg0 envcontext.ProviderCallContext) (network.ProviderAddresses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses", arg0)
	ret0, _ := ret[0].(network.ProviderAddresses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Addresses indicates an expected call of Addresses.
func (mr *MockInstanceMockRecorder) Addresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockInstance)(nil).Addresses), arg0)
}

// Id mocks base method.
func (m *MockInstance) Id() instance.Id {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(instance.Id)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockInstanceMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockInstance)(nil).Id))
}

// Status mocks base method.
func (m *MockInstance) Status(arg0 envcontext.ProviderCallContext) instance.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(instance.Status)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockInstanceMockRecorder) Status(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockInstance)(nil).Status), arg0)
}
