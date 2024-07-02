// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/charm/service (interfaces: State,WatcherFactory)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination package_mock_test.go github.com/juju/juju/domain/charm/service State,WatcherFactory
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	changestream "github.com/juju/juju/core/changestream"
	charm "github.com/juju/juju/core/charm"
	watcher "github.com/juju/juju/core/watcher"
	charm0 "github.com/juju/juju/domain/charm"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// GetCharmActions mocks base method.
func (m *MockState) GetCharmActions(arg0 context.Context, arg1 charm.ID) (charm0.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmActions", arg0, arg1)
	ret0, _ := ret[0].(charm0.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmActions indicates an expected call of GetCharmActions.
func (mr *MockStateMockRecorder) GetCharmActions(arg0, arg1 any) *MockStateGetCharmActionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmActions", reflect.TypeOf((*MockState)(nil).GetCharmActions), arg0, arg1)
	return &MockStateGetCharmActionsCall{Call: call}
}

// MockStateGetCharmActionsCall wrap *gomock.Call
type MockStateGetCharmActionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetCharmActionsCall) Return(arg0 charm0.Actions, arg1 error) *MockStateGetCharmActionsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetCharmActionsCall) Do(f func(context.Context, charm.ID) (charm0.Actions, error)) *MockStateGetCharmActionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetCharmActionsCall) DoAndReturn(f func(context.Context, charm.ID) (charm0.Actions, error)) *MockStateGetCharmActionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmConfig mocks base method.
func (m *MockState) GetCharmConfig(arg0 context.Context, arg1 charm.ID) (charm0.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmConfig", arg0, arg1)
	ret0, _ := ret[0].(charm0.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmConfig indicates an expected call of GetCharmConfig.
func (mr *MockStateMockRecorder) GetCharmConfig(arg0, arg1 any) *MockStateGetCharmConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmConfig", reflect.TypeOf((*MockState)(nil).GetCharmConfig), arg0, arg1)
	return &MockStateGetCharmConfigCall{Call: call}
}

// MockStateGetCharmConfigCall wrap *gomock.Call
type MockStateGetCharmConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetCharmConfigCall) Return(arg0 charm0.Config, arg1 error) *MockStateGetCharmConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetCharmConfigCall) Do(f func(context.Context, charm.ID) (charm0.Config, error)) *MockStateGetCharmConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetCharmConfigCall) DoAndReturn(f func(context.Context, charm.ID) (charm0.Config, error)) *MockStateGetCharmConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmIDByLatestRevision mocks base method.
func (m *MockState) GetCharmIDByLatestRevision(arg0 context.Context, arg1 string) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmIDByLatestRevision", arg0, arg1)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmIDByLatestRevision indicates an expected call of GetCharmIDByLatestRevision.
func (mr *MockStateMockRecorder) GetCharmIDByLatestRevision(arg0, arg1 any) *MockStateGetCharmIDByLatestRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmIDByLatestRevision", reflect.TypeOf((*MockState)(nil).GetCharmIDByLatestRevision), arg0, arg1)
	return &MockStateGetCharmIDByLatestRevisionCall{Call: call}
}

// MockStateGetCharmIDByLatestRevisionCall wrap *gomock.Call
type MockStateGetCharmIDByLatestRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetCharmIDByLatestRevisionCall) Return(arg0 charm.ID, arg1 error) *MockStateGetCharmIDByLatestRevisionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetCharmIDByLatestRevisionCall) Do(f func(context.Context, string) (charm.ID, error)) *MockStateGetCharmIDByLatestRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetCharmIDByLatestRevisionCall) DoAndReturn(f func(context.Context, string) (charm.ID, error)) *MockStateGetCharmIDByLatestRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmIDByRevision mocks base method.
func (m *MockState) GetCharmIDByRevision(arg0 context.Context, arg1 string, arg2 int) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmIDByRevision", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmIDByRevision indicates an expected call of GetCharmIDByRevision.
func (mr *MockStateMockRecorder) GetCharmIDByRevision(arg0, arg1, arg2 any) *MockStateGetCharmIDByRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmIDByRevision", reflect.TypeOf((*MockState)(nil).GetCharmIDByRevision), arg0, arg1, arg2)
	return &MockStateGetCharmIDByRevisionCall{Call: call}
}

// MockStateGetCharmIDByRevisionCall wrap *gomock.Call
type MockStateGetCharmIDByRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetCharmIDByRevisionCall) Return(arg0 charm.ID, arg1 error) *MockStateGetCharmIDByRevisionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetCharmIDByRevisionCall) Do(f func(context.Context, string, int) (charm.ID, error)) *MockStateGetCharmIDByRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetCharmIDByRevisionCall) DoAndReturn(f func(context.Context, string, int) (charm.ID, error)) *MockStateGetCharmIDByRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmLXDProfile mocks base method.
func (m *MockState) GetCharmLXDProfile(arg0 context.Context, arg1 charm.ID) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmLXDProfile", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmLXDProfile indicates an expected call of GetCharmLXDProfile.
func (mr *MockStateMockRecorder) GetCharmLXDProfile(arg0, arg1 any) *MockStateGetCharmLXDProfileCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmLXDProfile", reflect.TypeOf((*MockState)(nil).GetCharmLXDProfile), arg0, arg1)
	return &MockStateGetCharmLXDProfileCall{Call: call}
}

// MockStateGetCharmLXDProfileCall wrap *gomock.Call
type MockStateGetCharmLXDProfileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetCharmLXDProfileCall) Return(arg0 []byte, arg1 error) *MockStateGetCharmLXDProfileCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetCharmLXDProfileCall) Do(f func(context.Context, charm.ID) ([]byte, error)) *MockStateGetCharmLXDProfileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetCharmLXDProfileCall) DoAndReturn(f func(context.Context, charm.ID) ([]byte, error)) *MockStateGetCharmLXDProfileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmManifest mocks base method.
func (m *MockState) GetCharmManifest(arg0 context.Context, arg1 charm.ID) (charm0.Manifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmManifest", arg0, arg1)
	ret0, _ := ret[0].(charm0.Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmManifest indicates an expected call of GetCharmManifest.
func (mr *MockStateMockRecorder) GetCharmManifest(arg0, arg1 any) *MockStateGetCharmManifestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmManifest", reflect.TypeOf((*MockState)(nil).GetCharmManifest), arg0, arg1)
	return &MockStateGetCharmManifestCall{Call: call}
}

// MockStateGetCharmManifestCall wrap *gomock.Call
type MockStateGetCharmManifestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetCharmManifestCall) Return(arg0 charm0.Manifest, arg1 error) *MockStateGetCharmManifestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetCharmManifestCall) Do(f func(context.Context, charm.ID) (charm0.Manifest, error)) *MockStateGetCharmManifestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetCharmManifestCall) DoAndReturn(f func(context.Context, charm.ID) (charm0.Manifest, error)) *MockStateGetCharmManifestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmMetadata mocks base method.
func (m *MockState) GetCharmMetadata(arg0 context.Context, arg1 charm.ID) (charm0.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmMetadata", arg0, arg1)
	ret0, _ := ret[0].(charm0.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmMetadata indicates an expected call of GetCharmMetadata.
func (mr *MockStateMockRecorder) GetCharmMetadata(arg0, arg1 any) *MockStateGetCharmMetadataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmMetadata", reflect.TypeOf((*MockState)(nil).GetCharmMetadata), arg0, arg1)
	return &MockStateGetCharmMetadataCall{Call: call}
}

// MockStateGetCharmMetadataCall wrap *gomock.Call
type MockStateGetCharmMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetCharmMetadataCall) Return(arg0 charm0.Metadata, arg1 error) *MockStateGetCharmMetadataCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetCharmMetadataCall) Do(f func(context.Context, charm.ID) (charm0.Metadata, error)) *MockStateGetCharmMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetCharmMetadataCall) DoAndReturn(f func(context.Context, charm.ID) (charm0.Metadata, error)) *MockStateGetCharmMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsCharmAvailable mocks base method.
func (m *MockState) IsCharmAvailable(arg0 context.Context, arg1 charm.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCharmAvailable", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCharmAvailable indicates an expected call of IsCharmAvailable.
func (mr *MockStateMockRecorder) IsCharmAvailable(arg0, arg1 any) *MockStateIsCharmAvailableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCharmAvailable", reflect.TypeOf((*MockState)(nil).IsCharmAvailable), arg0, arg1)
	return &MockStateIsCharmAvailableCall{Call: call}
}

// MockStateIsCharmAvailableCall wrap *gomock.Call
type MockStateIsCharmAvailableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateIsCharmAvailableCall) Return(arg0 bool, arg1 error) *MockStateIsCharmAvailableCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateIsCharmAvailableCall) Do(f func(context.Context, charm.ID) (bool, error)) *MockStateIsCharmAvailableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateIsCharmAvailableCall) DoAndReturn(f func(context.Context, charm.ID) (bool, error)) *MockStateIsCharmAvailableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsControllerCharm mocks base method.
func (m *MockState) IsControllerCharm(arg0 context.Context, arg1 charm.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsControllerCharm", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsControllerCharm indicates an expected call of IsControllerCharm.
func (mr *MockStateMockRecorder) IsControllerCharm(arg0, arg1 any) *MockStateIsControllerCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsControllerCharm", reflect.TypeOf((*MockState)(nil).IsControllerCharm), arg0, arg1)
	return &MockStateIsControllerCharmCall{Call: call}
}

// MockStateIsControllerCharmCall wrap *gomock.Call
type MockStateIsControllerCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateIsControllerCharmCall) Return(arg0 bool, arg1 error) *MockStateIsControllerCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateIsControllerCharmCall) Do(f func(context.Context, charm.ID) (bool, error)) *MockStateIsControllerCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateIsControllerCharmCall) DoAndReturn(f func(context.Context, charm.ID) (bool, error)) *MockStateIsControllerCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsSubordinateCharm mocks base method.
func (m *MockState) IsSubordinateCharm(arg0 context.Context, arg1 charm.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSubordinateCharm", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSubordinateCharm indicates an expected call of IsSubordinateCharm.
func (mr *MockStateMockRecorder) IsSubordinateCharm(arg0, arg1 any) *MockStateIsSubordinateCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSubordinateCharm", reflect.TypeOf((*MockState)(nil).IsSubordinateCharm), arg0, arg1)
	return &MockStateIsSubordinateCharmCall{Call: call}
}

// MockStateIsSubordinateCharmCall wrap *gomock.Call
type MockStateIsSubordinateCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateIsSubordinateCharmCall) Return(arg0 bool, arg1 error) *MockStateIsSubordinateCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateIsSubordinateCharmCall) Do(f func(context.Context, charm.ID) (bool, error)) *MockStateIsSubordinateCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateIsSubordinateCharmCall) DoAndReturn(f func(context.Context, charm.ID) (bool, error)) *MockStateIsSubordinateCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReserveCharmRevision mocks base method.
func (m *MockState) ReserveCharmRevision(arg0 context.Context, arg1 charm.ID, arg2 int) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveCharmRevision", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReserveCharmRevision indicates an expected call of ReserveCharmRevision.
func (mr *MockStateMockRecorder) ReserveCharmRevision(arg0, arg1, arg2 any) *MockStateReserveCharmRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveCharmRevision", reflect.TypeOf((*MockState)(nil).ReserveCharmRevision), arg0, arg1, arg2)
	return &MockStateReserveCharmRevisionCall{Call: call}
}

// MockStateReserveCharmRevisionCall wrap *gomock.Call
type MockStateReserveCharmRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateReserveCharmRevisionCall) Return(arg0 charm.ID, arg1 error) *MockStateReserveCharmRevisionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateReserveCharmRevisionCall) Do(f func(context.Context, charm.ID, int) (charm.ID, error)) *MockStateReserveCharmRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateReserveCharmRevisionCall) DoAndReturn(f func(context.Context, charm.ID, int) (charm.ID, error)) *MockStateReserveCharmRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetCharm mocks base method.
func (m *MockState) SetCharm(arg0 context.Context, arg1 charm0.Charm) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharm", arg0, arg1)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCharm indicates an expected call of SetCharm.
func (mr *MockStateMockRecorder) SetCharm(arg0, arg1 any) *MockStateSetCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharm", reflect.TypeOf((*MockState)(nil).SetCharm), arg0, arg1)
	return &MockStateSetCharmCall{Call: call}
}

// MockStateSetCharmCall wrap *gomock.Call
type MockStateSetCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetCharmCall) Return(arg0 charm.ID, arg1 error) *MockStateSetCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetCharmCall) Do(f func(context.Context, charm0.Charm) (charm.ID, error)) *MockStateSetCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetCharmCall) DoAndReturn(f func(context.Context, charm0.Charm) (charm.ID, error)) *MockStateSetCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetCharmAvailable mocks base method.
func (m *MockState) SetCharmAvailable(arg0 context.Context, arg1 charm.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharmAvailable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCharmAvailable indicates an expected call of SetCharmAvailable.
func (mr *MockStateMockRecorder) SetCharmAvailable(arg0, arg1 any) *MockStateSetCharmAvailableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharmAvailable", reflect.TypeOf((*MockState)(nil).SetCharmAvailable), arg0, arg1)
	return &MockStateSetCharmAvailableCall{Call: call}
}

// MockStateSetCharmAvailableCall wrap *gomock.Call
type MockStateSetCharmAvailableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetCharmAvailableCall) Return(arg0 error) *MockStateSetCharmAvailableCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetCharmAvailableCall) Do(f func(context.Context, charm.ID) error) *MockStateSetCharmAvailableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetCharmAvailableCall) DoAndReturn(f func(context.Context, charm.ID) error) *MockStateSetCharmAvailableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SupportsContainers mocks base method.
func (m *MockState) SupportsContainers(arg0 context.Context, arg1 charm.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsContainers", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SupportsContainers indicates an expected call of SupportsContainers.
func (mr *MockStateMockRecorder) SupportsContainers(arg0, arg1 any) *MockStateSupportsContainersCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsContainers", reflect.TypeOf((*MockState)(nil).SupportsContainers), arg0, arg1)
	return &MockStateSupportsContainersCall{Call: call}
}

// MockStateSupportsContainersCall wrap *gomock.Call
type MockStateSupportsContainersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSupportsContainersCall) Return(arg0 bool, arg1 error) *MockStateSupportsContainersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSupportsContainersCall) Do(f func(context.Context, charm.ID) (bool, error)) *MockStateSupportsContainersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSupportsContainersCall) DoAndReturn(f func(context.Context, charm.ID) (bool, error)) *MockStateSupportsContainersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockWatcherFactory is a mock of WatcherFactory interface.
type MockWatcherFactory struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherFactoryMockRecorder
}

// MockWatcherFactoryMockRecorder is the mock recorder for MockWatcherFactory.
type MockWatcherFactoryMockRecorder struct {
	mock *MockWatcherFactory
}

// NewMockWatcherFactory creates a new mock instance.
func NewMockWatcherFactory(ctrl *gomock.Controller) *MockWatcherFactory {
	mock := &MockWatcherFactory{ctrl: ctrl}
	mock.recorder = &MockWatcherFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcherFactory) EXPECT() *MockWatcherFactoryMockRecorder {
	return m.recorder
}

// NewUUIDsWatcher mocks base method.
func (m *MockWatcherFactory) NewUUIDsWatcher(arg0 string, arg1 changestream.ChangeType) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUUIDsWatcher", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUUIDsWatcher indicates an expected call of NewUUIDsWatcher.
func (mr *MockWatcherFactoryMockRecorder) NewUUIDsWatcher(arg0, arg1 any) *MockWatcherFactoryNewUUIDsWatcherCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUUIDsWatcher", reflect.TypeOf((*MockWatcherFactory)(nil).NewUUIDsWatcher), arg0, arg1)
	return &MockWatcherFactoryNewUUIDsWatcherCall{Call: call}
}

// MockWatcherFactoryNewUUIDsWatcherCall wrap *gomock.Call
type MockWatcherFactoryNewUUIDsWatcherCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockWatcherFactoryNewUUIDsWatcherCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockWatcherFactoryNewUUIDsWatcherCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockWatcherFactoryNewUUIDsWatcherCall) Do(f func(string, changestream.ChangeType) (watcher.Watcher[[]string], error)) *MockWatcherFactoryNewUUIDsWatcherCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockWatcherFactoryNewUUIDsWatcherCall) DoAndReturn(f func(string, changestream.ChangeType) (watcher.Watcher[[]string], error)) *MockWatcherFactoryNewUUIDsWatcherCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
