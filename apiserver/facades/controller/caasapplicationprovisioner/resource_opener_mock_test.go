// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/resource (interfaces: Opener)
//
// Generated by this command:
//
//	mockgen -package caasapplicationprovisioner_test -destination resource_opener_mock_test.go github.com/juju/juju/core/resource Opener
//

// Package caasapplicationprovisioner_test is a generated GoMock package.
package caasapplicationprovisioner_test

import (
	context "context"
	reflect "reflect"

	resource "github.com/juju/juju/core/resource"
	gomock "go.uber.org/mock/gomock"
)

// MockOpener is a mock of Opener interface.
type MockOpener struct {
	ctrl     *gomock.Controller
	recorder *MockOpenerMockRecorder
}

// MockOpenerMockRecorder is the mock recorder for MockOpener.
type MockOpenerMockRecorder struct {
	mock *MockOpener
}

// NewMockOpener creates a new mock instance.
func NewMockOpener(ctrl *gomock.Controller) *MockOpener {
	mock := &MockOpener{ctrl: ctrl}
	mock.recorder = &MockOpenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpener) EXPECT() *MockOpenerMockRecorder {
	return m.recorder
}

// OpenResource mocks base method.
func (m *MockOpener) OpenResource(arg0 context.Context, arg1 string) (resource.Opened, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenResource", arg0, arg1)
	ret0, _ := ret[0].(resource.Opened)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenResource indicates an expected call of OpenResource.
func (mr *MockOpenerMockRecorder) OpenResource(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenResource", reflect.TypeOf((*MockOpener)(nil).OpenResource), arg0, arg1)
}

// SetResourceUsed mocks base method.
func (m *MockOpener) SetResourceUsed(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetResourceUsed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetResourceUsed indicates an expected call of SetResourceUsed.
func (mr *MockOpenerMockRecorder) SetResource(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResourceUsed", reflect.TypeOf((*MockOpener)(nil).SetResourceUsed), arg0, arg1)
}
