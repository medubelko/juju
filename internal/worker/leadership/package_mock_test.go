// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/leadership (interfaces: Claimer)
//
// Generated by this command:
//
//	mockgen -package leadership_test -destination package_mock_test.go github.com/juju/juju/core/leadership Claimer
//

// Package leadership_test is a generated GoMock package.
package leadership_test

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockClaimer is a mock of Claimer interface.
type MockClaimer struct {
	ctrl     *gomock.Controller
	recorder *MockClaimerMockRecorder
}

// MockClaimerMockRecorder is the mock recorder for MockClaimer.
type MockClaimerMockRecorder struct {
	mock *MockClaimer
}

// NewMockClaimer creates a new mock instance.
func NewMockClaimer(ctrl *gomock.Controller) *MockClaimer {
	mock := &MockClaimer{ctrl: ctrl}
	mock.recorder = &MockClaimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClaimer) EXPECT() *MockClaimerMockRecorder {
	return m.recorder
}

// BlockUntilLeadershipReleased mocks base method.
func (m *MockClaimer) BlockUntilLeadershipReleased(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockUntilLeadershipReleased", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BlockUntilLeadershipReleased indicates an expected call of BlockUntilLeadershipReleased.
func (mr *MockClaimerMockRecorder) BlockUntilLeadershipReleased(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockUntilLeadershipReleased", reflect.TypeOf((*MockClaimer)(nil).BlockUntilLeadershipReleased), arg0, arg1)
}

// ClaimLeadership mocks base method.
func (m *MockClaimer) ClaimLeadership(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimLeadership", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClaimLeadership indicates an expected call of ClaimLeadership.
func (mr *MockClaimerMockRecorder) ClaimLeadership(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimLeadership", reflect.TypeOf((*MockClaimer)(nil).ClaimLeadership), arg0, arg1, arg2, arg3)
}
