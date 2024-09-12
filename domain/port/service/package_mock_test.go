// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/port/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination package_mock_test.go github.com/juju/juju/domain/port/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	network "github.com/juju/juju/core/network"
	domain "github.com/juju/juju/domain"
	port "github.com/juju/juju/domain/port"
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

// GetApplicationOpenedPorts mocks base method.
func (m *MockState) GetApplicationOpenedPorts(arg0 context.Context, arg1 string) (port.UnitEndpointPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationOpenedPorts", arg0, arg1)
	ret0, _ := ret[0].(port.UnitEndpointPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationOpenedPorts indicates an expected call of GetApplicationOpenedPorts.
func (mr *MockStateMockRecorder) GetApplicationOpenedPorts(arg0, arg1 any) *MockStateGetApplicationOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationOpenedPorts", reflect.TypeOf((*MockState)(nil).GetApplicationOpenedPorts), arg0, arg1)
	return &MockStateGetApplicationOpenedPortsCall{Call: call}
}

// MockStateGetApplicationOpenedPortsCall wrap *gomock.Call
type MockStateGetApplicationOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetApplicationOpenedPortsCall) Return(arg0 port.UnitEndpointPortRanges, arg1 error) *MockStateGetApplicationOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetApplicationOpenedPortsCall) Do(f func(context.Context, string) (port.UnitEndpointPortRanges, error)) *MockStateGetApplicationOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetApplicationOpenedPortsCall) DoAndReturn(f func(context.Context, string) (port.UnitEndpointPortRanges, error)) *MockStateGetApplicationOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetColocatedOpenedPorts mocks base method.
func (m *MockState) GetColocatedOpenedPorts(arg0 domain.AtomicContext, arg1 string) ([]network.PortRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetColocatedOpenedPorts", arg0, arg1)
	ret0, _ := ret[0].([]network.PortRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetColocatedOpenedPorts indicates an expected call of GetColocatedOpenedPorts.
func (mr *MockStateMockRecorder) GetColocatedOpenedPorts(arg0, arg1 any) *MockStateGetColocatedOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColocatedOpenedPorts", reflect.TypeOf((*MockState)(nil).GetColocatedOpenedPorts), arg0, arg1)
	return &MockStateGetColocatedOpenedPortsCall{Call: call}
}

// MockStateGetColocatedOpenedPortsCall wrap *gomock.Call
type MockStateGetColocatedOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetColocatedOpenedPortsCall) Return(arg0 []network.PortRange, arg1 error) *MockStateGetColocatedOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetColocatedOpenedPortsCall) Do(f func(domain.AtomicContext, string) ([]network.PortRange, error)) *MockStateGetColocatedOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetColocatedOpenedPortsCall) DoAndReturn(f func(domain.AtomicContext, string) ([]network.PortRange, error)) *MockStateGetColocatedOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetEndpoints mocks base method.
func (m *MockState) GetEndpoints(arg0 domain.AtomicContext, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEndpoints indicates an expected call of GetEndpoints.
func (mr *MockStateMockRecorder) GetEndpoints(arg0, arg1 any) *MockStateGetEndpointsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndpoints", reflect.TypeOf((*MockState)(nil).GetEndpoints), arg0, arg1)
	return &MockStateGetEndpointsCall{Call: call}
}

// MockStateGetEndpointsCall wrap *gomock.Call
type MockStateGetEndpointsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetEndpointsCall) Return(arg0 []string, arg1 error) *MockStateGetEndpointsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetEndpointsCall) Do(f func(domain.AtomicContext, string) ([]string, error)) *MockStateGetEndpointsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetEndpointsCall) DoAndReturn(f func(domain.AtomicContext, string) ([]string, error)) *MockStateGetEndpointsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMachineOpenedPorts mocks base method.
func (m *MockState) GetMachineOpenedPorts(arg0 context.Context, arg1 string) (map[string]network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineOpenedPorts", arg0, arg1)
	ret0, _ := ret[0].(map[string]network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineOpenedPorts indicates an expected call of GetMachineOpenedPorts.
func (mr *MockStateMockRecorder) GetMachineOpenedPorts(arg0, arg1 any) *MockStateGetMachineOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineOpenedPorts", reflect.TypeOf((*MockState)(nil).GetMachineOpenedPorts), arg0, arg1)
	return &MockStateGetMachineOpenedPortsCall{Call: call}
}

// MockStateGetMachineOpenedPortsCall wrap *gomock.Call
type MockStateGetMachineOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetMachineOpenedPortsCall) Return(arg0 map[string]network.GroupedPortRanges, arg1 error) *MockStateGetMachineOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetMachineOpenedPortsCall) Do(f func(context.Context, string) (map[string]network.GroupedPortRanges, error)) *MockStateGetMachineOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetMachineOpenedPortsCall) DoAndReturn(f func(context.Context, string) (map[string]network.GroupedPortRanges, error)) *MockStateGetMachineOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOpenedEndpointPorts mocks base method.
func (m *MockState) GetOpenedEndpointPorts(arg0 domain.AtomicContext, arg1, arg2 string) ([]network.PortRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenedEndpointPorts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]network.PortRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenedEndpointPorts indicates an expected call of GetOpenedEndpointPorts.
func (mr *MockStateMockRecorder) GetOpenedEndpointPorts(arg0, arg1, arg2 any) *MockStateGetOpenedEndpointPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenedEndpointPorts", reflect.TypeOf((*MockState)(nil).GetOpenedEndpointPorts), arg0, arg1, arg2)
	return &MockStateGetOpenedEndpointPortsCall{Call: call}
}

// MockStateGetOpenedEndpointPortsCall wrap *gomock.Call
type MockStateGetOpenedEndpointPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetOpenedEndpointPortsCall) Return(arg0 []network.PortRange, arg1 error) *MockStateGetOpenedEndpointPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetOpenedEndpointPortsCall) Do(f func(domain.AtomicContext, string, string) ([]network.PortRange, error)) *MockStateGetOpenedEndpointPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetOpenedEndpointPortsCall) DoAndReturn(f func(domain.AtomicContext, string, string) ([]network.PortRange, error)) *MockStateGetOpenedEndpointPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitOpenedPorts mocks base method.
func (m *MockState) GetUnitOpenedPorts(arg0 context.Context, arg1 string) (network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitOpenedPorts", arg0, arg1)
	ret0, _ := ret[0].(network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitOpenedPorts indicates an expected call of GetUnitOpenedPorts.
func (mr *MockStateMockRecorder) GetUnitOpenedPorts(arg0, arg1 any) *MockStateGetUnitOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitOpenedPorts", reflect.TypeOf((*MockState)(nil).GetUnitOpenedPorts), arg0, arg1)
	return &MockStateGetUnitOpenedPortsCall{Call: call}
}

// MockStateGetUnitOpenedPortsCall wrap *gomock.Call
type MockStateGetUnitOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetUnitOpenedPortsCall) Return(arg0 network.GroupedPortRanges, arg1 error) *MockStateGetUnitOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetUnitOpenedPortsCall) Do(f func(context.Context, string) (network.GroupedPortRanges, error)) *MockStateGetUnitOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetUnitOpenedPortsCall) DoAndReturn(f func(context.Context, string) (network.GroupedPortRanges, error)) *MockStateGetUnitOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RunAtomic mocks base method.
func (m *MockState) RunAtomic(arg0 context.Context, arg1 func(domain.AtomicContext) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunAtomic", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunAtomic indicates an expected call of RunAtomic.
func (mr *MockStateMockRecorder) RunAtomic(arg0, arg1 any) *MockStateRunAtomicCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAtomic", reflect.TypeOf((*MockState)(nil).RunAtomic), arg0, arg1)
	return &MockStateRunAtomicCall{Call: call}
}

// MockStateRunAtomicCall wrap *gomock.Call
type MockStateRunAtomicCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateRunAtomicCall) Return(arg0 error) *MockStateRunAtomicCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateRunAtomicCall) Do(f func(context.Context, func(domain.AtomicContext) error) error) *MockStateRunAtomicCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateRunAtomicCall) DoAndReturn(f func(context.Context, func(domain.AtomicContext) error) error) *MockStateRunAtomicCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUnitPorts mocks base method.
func (m *MockState) UpdateUnitPorts(arg0 domain.AtomicContext, arg1 string, arg2, arg3 network.GroupedPortRanges) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnitPorts", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUnitPorts indicates an expected call of UpdateUnitPorts.
func (mr *MockStateMockRecorder) UpdateUnitPorts(arg0, arg1, arg2, arg3 any) *MockStateUpdateUnitPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnitPorts", reflect.TypeOf((*MockState)(nil).UpdateUnitPorts), arg0, arg1, arg2, arg3)
	return &MockStateUpdateUnitPortsCall{Call: call}
}

// MockStateUpdateUnitPortsCall wrap *gomock.Call
type MockStateUpdateUnitPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateUpdateUnitPortsCall) Return(arg0 error) *MockStateUpdateUnitPortsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateUpdateUnitPortsCall) Do(f func(domain.AtomicContext, string, network.GroupedPortRanges, network.GroupedPortRanges) error) *MockStateUpdateUnitPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateUpdateUnitPortsCall) DoAndReturn(f func(domain.AtomicContext, string, network.GroupedPortRanges, network.GroupedPortRanges) error) *MockStateUpdateUnitPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
