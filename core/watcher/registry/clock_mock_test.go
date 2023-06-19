// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/clock (interfaces: Clock)

// Package registry is a generated GoMock package.
package registry

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	clock "github.com/juju/clock"
)

// MockClock is a mock of Clock interface.
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock.
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance.
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// After mocks base method.
func (m *MockClock) After(arg0 time.Duration) <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "After", arg0)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// After indicates an expected call of After.
func (mr *MockClockMockRecorder) After(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "After", reflect.TypeOf((*MockClock)(nil).After), arg0)
}

// AfterFunc mocks base method.
func (m *MockClock) AfterFunc(arg0 time.Duration, arg1 func()) clock.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterFunc", arg0, arg1)
	ret0, _ := ret[0].(clock.Timer)
	return ret0
}

// AfterFunc indicates an expected call of AfterFunc.
func (mr *MockClockMockRecorder) AfterFunc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFunc", reflect.TypeOf((*MockClock)(nil).AfterFunc), arg0, arg1)
}

// NewTimer mocks base method.
func (m *MockClock) NewTimer(arg0 time.Duration) clock.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTimer", arg0)
	ret0, _ := ret[0].(clock.Timer)
	return ret0
}

// NewTimer indicates an expected call of NewTimer.
func (mr *MockClockMockRecorder) NewTimer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTimer", reflect.TypeOf((*MockClock)(nil).NewTimer), arg0)
}

// Now mocks base method.
func (m *MockClock) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockClockMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
}
