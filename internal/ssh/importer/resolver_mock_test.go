// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/ssh/importer (interfaces: Resolver)
//
// Generated by this command:
//
//	mockgen -typed -package importer -destination resolver_mock_test.go github.com/juju/juju/internal/ssh/importer Resolver
//

// Package importer is a generated GoMock package.
package importer

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockResolver is a mock of Resolver interface.
type MockResolver struct {
	ctrl     *gomock.Controller
	recorder *MockResolverMockRecorder
}

// MockResolverMockRecorder is the mock recorder for MockResolver.
type MockResolverMockRecorder struct {
	mock *MockResolver
}

// NewMockResolver creates a new mock instance.
func NewMockResolver(ctrl *gomock.Controller) *MockResolver {
	mock := &MockResolver{ctrl: ctrl}
	mock.recorder = &MockResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResolver) EXPECT() *MockResolverMockRecorder {
	return m.recorder
}

// PublicKeysForSubject mocks base method.
func (m *MockResolver) PublicKeysForSubject(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicKeysForSubject", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicKeysForSubject indicates an expected call of PublicKeysForSubject.
func (mr *MockResolverMockRecorder) PublicKeysForSubject(arg0, arg1 any) *MockResolverPublicKeysForSubjectCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicKeysForSubject", reflect.TypeOf((*MockResolver)(nil).PublicKeysForSubject), arg0, arg1)
	return &MockResolverPublicKeysForSubjectCall{Call: call}
}

// MockResolverPublicKeysForSubjectCall wrap *gomock.Call
type MockResolverPublicKeysForSubjectCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResolverPublicKeysForSubjectCall) Return(arg0 []string, arg1 error) *MockResolverPublicKeysForSubjectCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResolverPublicKeysForSubjectCall) Do(f func(context.Context, string) ([]string, error)) *MockResolverPublicKeysForSubjectCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResolverPublicKeysForSubjectCall) DoAndReturn(f func(context.Context, string) ([]string, error)) *MockResolverPublicKeysForSubjectCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
