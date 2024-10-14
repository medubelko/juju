// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/firewaller (interfaces: ControllerConfigService,ModelConfigService,NetworkService,ApplicationService,MachineService,PortService)
//
// Generated by this command:
//
//	mockgen -typed -package firewaller_test -destination service_mock_test.go github.com/juju/juju/apiserver/facades/controller/firewaller ControllerConfigService,ModelConfigService,NetworkService,ApplicationService,MachineService,PortService
//

// Package firewaller_test is a generated GoMock package.
package firewaller_test

import (
	context "context"
	reflect "reflect"

	set "github.com/juju/collections/set"
	controller "github.com/juju/juju/controller"
	instance "github.com/juju/juju/core/instance"
	life "github.com/juju/juju/core/life"
	machine "github.com/juju/juju/core/machine"
	network "github.com/juju/juju/core/network"
	watcher "github.com/juju/juju/core/watcher"
	port "github.com/juju/juju/domain/port"
	config "github.com/juju/juju/environs/config"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 any) *MockControllerConfigServiceControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
	return &MockControllerConfigServiceControllerConfigCall{Call: call}
}

// MockControllerConfigServiceControllerConfigCall wrap *gomock.Call
type MockControllerConfigServiceControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerConfigServiceControllerConfigCall) Return(arg0 controller.Config, arg1 error) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerConfigServiceControllerConfigCall) Do(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerConfigServiceControllerConfigCall) DoAndReturn(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelConfigService is a mock of ModelConfigService interface.
type MockModelConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigServiceMockRecorder
}

// MockModelConfigServiceMockRecorder is the mock recorder for MockModelConfigService.
type MockModelConfigServiceMockRecorder struct {
	mock *MockModelConfigService
}

// NewMockModelConfigService creates a new mock instance.
func NewMockModelConfigService(ctrl *gomock.Controller) *MockModelConfigService {
	mock := &MockModelConfigService{ctrl: ctrl}
	mock.recorder = &MockModelConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigService) EXPECT() *MockModelConfigServiceMockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockModelConfigService) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelConfigServiceMockRecorder) ModelConfig(arg0 any) *MockModelConfigServiceModelConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModelConfigService)(nil).ModelConfig), arg0)
	return &MockModelConfigServiceModelConfigCall{Call: call}
}

// MockModelConfigServiceModelConfigCall wrap *gomock.Call
type MockModelConfigServiceModelConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigServiceModelConfigCall) Return(arg0 *config.Config, arg1 error) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigServiceModelConfigCall) Do(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigServiceModelConfigCall) DoAndReturn(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Watch mocks base method.
func (m *MockModelConfigService) Watch() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockModelConfigServiceMockRecorder) Watch() *MockModelConfigServiceWatchCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockModelConfigService)(nil).Watch))
	return &MockModelConfigServiceWatchCall{Call: call}
}

// MockModelConfigServiceWatchCall wrap *gomock.Call
type MockModelConfigServiceWatchCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigServiceWatchCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockModelConfigServiceWatchCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigServiceWatchCall) Do(f func() (watcher.Watcher[[]string], error)) *MockModelConfigServiceWatchCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigServiceWatchCall) DoAndReturn(f func() (watcher.Watcher[[]string], error)) *MockModelConfigServiceWatchCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// GetAllSpaces mocks base method.
func (m *MockNetworkService) GetAllSpaces(arg0 context.Context) (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpaces", arg0)
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpaces indicates an expected call of GetAllSpaces.
func (mr *MockNetworkServiceMockRecorder) GetAllSpaces(arg0 any) *MockNetworkServiceGetAllSpacesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpaces", reflect.TypeOf((*MockNetworkService)(nil).GetAllSpaces), arg0)
	return &MockNetworkServiceGetAllSpacesCall{Call: call}
}

// MockNetworkServiceGetAllSpacesCall wrap *gomock.Call
type MockNetworkServiceGetAllSpacesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceGetAllSpacesCall) Return(arg0 network.SpaceInfos, arg1 error) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceGetAllSpacesCall) Do(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceGetAllSpacesCall) DoAndReturn(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchSubnets mocks base method.
func (m *MockNetworkService) WatchSubnets(arg0 context.Context, arg1 set.Strings) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchSubnets", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchSubnets indicates an expected call of WatchSubnets.
func (mr *MockNetworkServiceMockRecorder) WatchSubnets(arg0, arg1 any) *MockNetworkServiceWatchSubnetsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchSubnets", reflect.TypeOf((*MockNetworkService)(nil).WatchSubnets), arg0, arg1)
	return &MockNetworkServiceWatchSubnetsCall{Call: call}
}

// MockNetworkServiceWatchSubnetsCall wrap *gomock.Call
type MockNetworkServiceWatchSubnetsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceWatchSubnetsCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockNetworkServiceWatchSubnetsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceWatchSubnetsCall) Do(f func(context.Context, set.Strings) (watcher.Watcher[[]string], error)) *MockNetworkServiceWatchSubnetsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceWatchSubnetsCall) DoAndReturn(f func(context.Context, set.Strings) (watcher.Watcher[[]string], error)) *MockNetworkServiceWatchSubnetsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// GetUnitLife mocks base method.
func (m *MockApplicationService) GetUnitLife(arg0 context.Context, arg1 string) (life.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitLife", arg0, arg1)
	ret0, _ := ret[0].(life.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitLife indicates an expected call of GetUnitLife.
func (mr *MockApplicationServiceMockRecorder) GetUnitLife(arg0, arg1 any) *MockApplicationServiceGetUnitLifeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitLife", reflect.TypeOf((*MockApplicationService)(nil).GetUnitLife), arg0, arg1)
	return &MockApplicationServiceGetUnitLifeCall{Call: call}
}

// MockApplicationServiceGetUnitLifeCall wrap *gomock.Call
type MockApplicationServiceGetUnitLifeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetUnitLifeCall) Return(arg0 life.Value, arg1 error) *MockApplicationServiceGetUnitLifeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetUnitLifeCall) Do(f func(context.Context, string) (life.Value, error)) *MockApplicationServiceGetUnitLifeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetUnitLifeCall) DoAndReturn(f func(context.Context, string) (life.Value, error)) *MockApplicationServiceGetUnitLifeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockMachineService is a mock of MachineService interface.
type MockMachineService struct {
	ctrl     *gomock.Controller
	recorder *MockMachineServiceMockRecorder
}

// MockMachineServiceMockRecorder is the mock recorder for MockMachineService.
type MockMachineServiceMockRecorder struct {
	mock *MockMachineService
}

// NewMockMachineService creates a new mock instance.
func NewMockMachineService(ctrl *gomock.Controller) *MockMachineService {
	mock := &MockMachineService{ctrl: ctrl}
	mock.recorder = &MockMachineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineService) EXPECT() *MockMachineServiceMockRecorder {
	return m.recorder
}

// EnsureDeadMachine mocks base method.
func (m *MockMachineService) EnsureDeadMachine(arg0 context.Context, arg1 machine.Name) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureDeadMachine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureDeadMachine indicates an expected call of EnsureDeadMachine.
func (mr *MockMachineServiceMockRecorder) EnsureDeadMachine(arg0, arg1 any) *MockMachineServiceEnsureDeadMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureDeadMachine", reflect.TypeOf((*MockMachineService)(nil).EnsureDeadMachine), arg0, arg1)
	return &MockMachineServiceEnsureDeadMachineCall{Call: call}
}

// MockMachineServiceEnsureDeadMachineCall wrap *gomock.Call
type MockMachineServiceEnsureDeadMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceEnsureDeadMachineCall) Return(arg0 error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceEnsureDeadMachineCall) Do(f func(context.Context, machine.Name) error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceEnsureDeadMachineCall) DoAndReturn(f func(context.Context, machine.Name) error) *MockMachineServiceEnsureDeadMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMachineUUID mocks base method.
func (m *MockMachineService) GetMachineUUID(arg0 context.Context, arg1 machine.Name) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineUUID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineUUID indicates an expected call of GetMachineUUID.
func (mr *MockMachineServiceMockRecorder) GetMachineUUID(arg0, arg1 any) *MockMachineServiceGetMachineUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineUUID", reflect.TypeOf((*MockMachineService)(nil).GetMachineUUID), arg0, arg1)
	return &MockMachineServiceGetMachineUUIDCall{Call: call}
}

// MockMachineServiceGetMachineUUIDCall wrap *gomock.Call
type MockMachineServiceGetMachineUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceGetMachineUUIDCall) Return(arg0 string, arg1 error) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceGetMachineUUIDCall) Do(f func(context.Context, machine.Name) (string, error)) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceGetMachineUUIDCall) DoAndReturn(f func(context.Context, machine.Name) (string, error)) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HardwareCharacteristics mocks base method.
func (m *MockMachineService) HardwareCharacteristics(arg0 context.Context, arg1 string) (*instance.HardwareCharacteristics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardwareCharacteristics", arg0, arg1)
	ret0, _ := ret[0].(*instance.HardwareCharacteristics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HardwareCharacteristics indicates an expected call of HardwareCharacteristics.
func (mr *MockMachineServiceMockRecorder) HardwareCharacteristics(arg0, arg1 any) *MockMachineServiceHardwareCharacteristicsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardwareCharacteristics", reflect.TypeOf((*MockMachineService)(nil).HardwareCharacteristics), arg0, arg1)
	return &MockMachineServiceHardwareCharacteristicsCall{Call: call}
}

// MockMachineServiceHardwareCharacteristicsCall wrap *gomock.Call
type MockMachineServiceHardwareCharacteristicsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceHardwareCharacteristicsCall) Return(arg0 *instance.HardwareCharacteristics, arg1 error) *MockMachineServiceHardwareCharacteristicsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceHardwareCharacteristicsCall) Do(f func(context.Context, string) (*instance.HardwareCharacteristics, error)) *MockMachineServiceHardwareCharacteristicsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceHardwareCharacteristicsCall) DoAndReturn(f func(context.Context, string) (*instance.HardwareCharacteristics, error)) *MockMachineServiceHardwareCharacteristicsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InstanceID mocks base method.
func (m *MockMachineService) InstanceID(arg0 context.Context, arg1 string) (instance.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceID", arg0, arg1)
	ret0, _ := ret[0].(instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceID indicates an expected call of InstanceID.
func (mr *MockMachineServiceMockRecorder) InstanceID(arg0, arg1 any) *MockMachineServiceInstanceIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockMachineService)(nil).InstanceID), arg0, arg1)
	return &MockMachineServiceInstanceIDCall{Call: call}
}

// MockMachineServiceInstanceIDCall wrap *gomock.Call
type MockMachineServiceInstanceIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceInstanceIDCall) Return(arg0 instance.Id, arg1 error) *MockMachineServiceInstanceIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceInstanceIDCall) Do(f func(context.Context, string) (instance.Id, error)) *MockMachineServiceInstanceIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceInstanceIDCall) DoAndReturn(f func(context.Context, string) (instance.Id, error)) *MockMachineServiceInstanceIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InstanceIDAndName mocks base method.
func (m *MockMachineService) InstanceIDAndName(arg0 context.Context, arg1 string) (instance.Id, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceIDAndName", arg0, arg1)
	ret0, _ := ret[0].(instance.Id)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// InstanceIDAndName indicates an expected call of InstanceIDAndName.
func (mr *MockMachineServiceMockRecorder) InstanceIDAndName(arg0, arg1 any) *MockMachineServiceInstanceIDAndNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceIDAndName", reflect.TypeOf((*MockMachineService)(nil).InstanceIDAndName), arg0, arg1)
	return &MockMachineServiceInstanceIDAndNameCall{Call: call}
}

// MockMachineServiceInstanceIDAndNameCall wrap *gomock.Call
type MockMachineServiceInstanceIDAndNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceInstanceIDAndNameCall) Return(arg0 instance.Id, arg1 string, arg2 error) *MockMachineServiceInstanceIDAndNameCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceInstanceIDAndNameCall) Do(f func(context.Context, string) (instance.Id, string, error)) *MockMachineServiceInstanceIDAndNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceInstanceIDAndNameCall) DoAndReturn(f func(context.Context, string) (instance.Id, string, error)) *MockMachineServiceInstanceIDAndNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockPortService is a mock of PortService interface.
type MockPortService struct {
	ctrl     *gomock.Controller
	recorder *MockPortServiceMockRecorder
}

// MockPortServiceMockRecorder is the mock recorder for MockPortService.
type MockPortServiceMockRecorder struct {
	mock *MockPortService
}

// NewMockPortService creates a new mock instance.
func NewMockPortService(ctrl *gomock.Controller) *MockPortService {
	mock := &MockPortService{ctrl: ctrl}
	mock.recorder = &MockPortServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortService) EXPECT() *MockPortServiceMockRecorder {
	return m.recorder
}

// GetMachineOpenedPortsAndSubnets mocks base method.
func (m *MockPortService) GetMachineOpenedPortsAndSubnets(arg0 context.Context, arg1 string) (map[string]port.GroupedPortRangesOnSubnets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineOpenedPortsAndSubnets", arg0, arg1)
	ret0, _ := ret[0].(map[string]port.GroupedPortRangesOnSubnets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineOpenedPortsAndSubnets indicates an expected call of GetMachineOpenedPortsAndSubnets.
func (mr *MockPortServiceMockRecorder) GetMachineOpenedPortsAndSubnets(arg0, arg1 any) *MockPortServiceGetMachineOpenedPortsAndSubnetsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineOpenedPortsAndSubnets", reflect.TypeOf((*MockPortService)(nil).GetMachineOpenedPortsAndSubnets), arg0, arg1)
	return &MockPortServiceGetMachineOpenedPortsAndSubnetsCall{Call: call}
}

// MockPortServiceGetMachineOpenedPortsAndSubnetsCall wrap *gomock.Call
type MockPortServiceGetMachineOpenedPortsAndSubnetsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPortServiceGetMachineOpenedPortsAndSubnetsCall) Return(arg0 map[string]port.GroupedPortRangesOnSubnets, arg1 error) *MockPortServiceGetMachineOpenedPortsAndSubnetsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPortServiceGetMachineOpenedPortsAndSubnetsCall) Do(f func(context.Context, string) (map[string]port.GroupedPortRangesOnSubnets, error)) *MockPortServiceGetMachineOpenedPortsAndSubnetsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPortServiceGetMachineOpenedPortsAndSubnetsCall) DoAndReturn(f func(context.Context, string) (map[string]port.GroupedPortRangesOnSubnets, error)) *MockPortServiceGetMachineOpenedPortsAndSubnetsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WatchOpenedPorts mocks base method.
func (m *MockPortService) WatchOpenedPorts(arg0 context.Context) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOpenedPorts", arg0)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchOpenedPorts indicates an expected call of WatchOpenedPorts.
func (mr *MockPortServiceMockRecorder) WatchOpenedPorts(arg0 any) *MockPortServiceWatchOpenedPortsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOpenedPorts", reflect.TypeOf((*MockPortService)(nil).WatchOpenedPorts), arg0)
	return &MockPortServiceWatchOpenedPortsCall{Call: call}
}

// MockPortServiceWatchOpenedPortsCall wrap *gomock.Call
type MockPortServiceWatchOpenedPortsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPortServiceWatchOpenedPortsCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockPortServiceWatchOpenedPortsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPortServiceWatchOpenedPortsCall) Do(f func(context.Context) (watcher.Watcher[[]string], error)) *MockPortServiceWatchOpenedPortsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPortServiceWatchOpenedPortsCall) DoAndReturn(f func(context.Context) (watcher.Watcher[[]string], error)) *MockPortServiceWatchOpenedPortsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
