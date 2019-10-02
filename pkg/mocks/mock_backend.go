// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/docker/stacks/pkg/interfaces (interfaces: BackendClient)

package mocks

import (
	types "github.com/docker/docker/api/types"
	events "github.com/docker/docker/api/types/events"
	filters "github.com/docker/docker/api/types/filters"
	swarm "github.com/docker/docker/api/types/swarm"
	interfaces "github.com/docker/stacks/pkg/interfaces"
	types0 "github.com/docker/stacks/pkg/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockBackendClient is a mock of BackendClient interface
type MockBackendClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackendClientMockRecorder
}

// MockBackendClientMockRecorder is the mock recorder for MockBackendClient
type MockBackendClientMockRecorder struct {
	mock *MockBackendClient
}

// NewMockBackendClient creates a new mock instance
func NewMockBackendClient(ctrl *gomock.Controller) *MockBackendClient {
	mock := &MockBackendClient{ctrl: ctrl}
	mock.recorder = &MockBackendClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockBackendClient) EXPECT() *MockBackendClientMockRecorder {
	return _m.recorder
}

// CreateConfig mocks base method
func (_m *MockBackendClient) CreateConfig(_param0 swarm.ConfigSpec) (string, error) {
	ret := _m.ctrl.Call(_m, "CreateConfig", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConfig indicates an expected call of CreateConfig
func (_mr *MockBackendClientMockRecorder) CreateConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateConfig", reflect.TypeOf((*MockBackendClient)(nil).CreateConfig), arg0)
}

// CreateNetwork mocks base method
func (_m *MockBackendClient) CreateNetwork(_param0 types.NetworkCreateRequest) (string, error) {
	ret := _m.ctrl.Call(_m, "CreateNetwork", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNetwork indicates an expected call of CreateNetwork
func (_mr *MockBackendClientMockRecorder) CreateNetwork(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateNetwork", reflect.TypeOf((*MockBackendClient)(nil).CreateNetwork), arg0)
}

// CreateSecret mocks base method
func (_m *MockBackendClient) CreateSecret(_param0 swarm.SecretSpec) (string, error) {
	ret := _m.ctrl.Call(_m, "CreateSecret", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret
func (_mr *MockBackendClientMockRecorder) CreateSecret(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateSecret", reflect.TypeOf((*MockBackendClient)(nil).CreateSecret), arg0)
}

// CreateService mocks base method
func (_m *MockBackendClient) CreateService(_param0 swarm.ServiceSpec, _param1 string, _param2 bool) (*types.ServiceCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateService", _param0, _param1, _param2)
	ret0, _ := ret[0].(*types.ServiceCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService
func (_mr *MockBackendClientMockRecorder) CreateService(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateService", reflect.TypeOf((*MockBackendClient)(nil).CreateService), arg0, arg1, arg2)
}

// CreateStack mocks base method
func (_m *MockBackendClient) CreateStack(_param0 types0.StackSpec) (types0.StackCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateStack", _param0)
	ret0, _ := ret[0].(types0.StackCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStack indicates an expected call of CreateStack
func (_mr *MockBackendClientMockRecorder) CreateStack(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateStack", reflect.TypeOf((*MockBackendClient)(nil).CreateStack), arg0)
}

// DeleteStack mocks base method
func (_m *MockBackendClient) DeleteStack(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteStack", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStack indicates an expected call of DeleteStack
func (_mr *MockBackendClientMockRecorder) DeleteStack(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteStack", reflect.TypeOf((*MockBackendClient)(nil).DeleteStack), arg0)
}

// GetConfig mocks base method
func (_m *MockBackendClient) GetConfig(_param0 string) (swarm.Config, error) {
	ret := _m.ctrl.Call(_m, "GetConfig", _param0)
	ret0, _ := ret[0].(swarm.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (_mr *MockBackendClientMockRecorder) GetConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetConfig", reflect.TypeOf((*MockBackendClient)(nil).GetConfig), arg0)
}

// GetConfigs mocks base method
func (_m *MockBackendClient) GetConfigs(_param0 types.ConfigListOptions) ([]swarm.Config, error) {
	ret := _m.ctrl.Call(_m, "GetConfigs", _param0)
	ret0, _ := ret[0].([]swarm.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigs indicates an expected call of GetConfigs
func (_mr *MockBackendClientMockRecorder) GetConfigs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetConfigs", reflect.TypeOf((*MockBackendClient)(nil).GetConfigs), arg0)
}

// GetNetwork mocks base method
func (_m *MockBackendClient) GetNetwork(_param0 string) (types.NetworkResource, error) {
	ret := _m.ctrl.Call(_m, "GetNetwork", _param0)
	ret0, _ := ret[0].(types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetwork indicates an expected call of GetNetwork
func (_mr *MockBackendClientMockRecorder) GetNetwork(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetNetwork", reflect.TypeOf((*MockBackendClient)(nil).GetNetwork), arg0)
}

// GetNetworks mocks base method
func (_m *MockBackendClient) GetNetworks(_param0 filters.Args) ([]types.NetworkResource, error) {
	ret := _m.ctrl.Call(_m, "GetNetworks", _param0)
	ret0, _ := ret[0].([]types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworks indicates an expected call of GetNetworks
func (_mr *MockBackendClientMockRecorder) GetNetworks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetNetworks", reflect.TypeOf((*MockBackendClient)(nil).GetNetworks), arg0)
}

// GetNetworksByName mocks base method
func (_m *MockBackendClient) GetNetworksByName(_param0 string) ([]types.NetworkResource, error) {
	ret := _m.ctrl.Call(_m, "GetNetworksByName", _param0)
	ret0, _ := ret[0].([]types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworksByName indicates an expected call of GetNetworksByName
func (_mr *MockBackendClientMockRecorder) GetNetworksByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetNetworksByName", reflect.TypeOf((*MockBackendClient)(nil).GetNetworksByName), arg0)
}

// GetNode mocks base method
func (_m *MockBackendClient) GetNode(_param0 string) (swarm.Node, error) {
	ret := _m.ctrl.Call(_m, "GetNode", _param0)
	ret0, _ := ret[0].(swarm.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNode indicates an expected call of GetNode
func (_mr *MockBackendClientMockRecorder) GetNode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetNode", reflect.TypeOf((*MockBackendClient)(nil).GetNode), arg0)
}

// GetSecret mocks base method
func (_m *MockBackendClient) GetSecret(_param0 string) (swarm.Secret, error) {
	ret := _m.ctrl.Call(_m, "GetSecret", _param0)
	ret0, _ := ret[0].(swarm.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (_mr *MockBackendClientMockRecorder) GetSecret(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSecret", reflect.TypeOf((*MockBackendClient)(nil).GetSecret), arg0)
}

// GetSecrets mocks base method
func (_m *MockBackendClient) GetSecrets(_param0 types.SecretListOptions) ([]swarm.Secret, error) {
	ret := _m.ctrl.Call(_m, "GetSecrets", _param0)
	ret0, _ := ret[0].([]swarm.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecrets indicates an expected call of GetSecrets
func (_mr *MockBackendClientMockRecorder) GetSecrets(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSecrets", reflect.TypeOf((*MockBackendClient)(nil).GetSecrets), arg0)
}

// GetService mocks base method
func (_m *MockBackendClient) GetService(_param0 string, _param1 bool) (swarm.Service, error) {
	ret := _m.ctrl.Call(_m, "GetService", _param0, _param1)
	ret0, _ := ret[0].(swarm.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (_mr *MockBackendClientMockRecorder) GetService(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetService", reflect.TypeOf((*MockBackendClient)(nil).GetService), arg0, arg1)
}

// GetServices mocks base method
func (_m *MockBackendClient) GetServices(_param0 types.ServiceListOptions) ([]swarm.Service, error) {
	ret := _m.ctrl.Call(_m, "GetServices", _param0)
	ret0, _ := ret[0].([]swarm.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServices indicates an expected call of GetServices
func (_mr *MockBackendClientMockRecorder) GetServices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetServices", reflect.TypeOf((*MockBackendClient)(nil).GetServices), arg0)
}

// GetSnapshotStack mocks base method
func (_m *MockBackendClient) GetSnapshotStack(_param0 string) (interfaces.SnapshotStack, error) {
	ret := _m.ctrl.Call(_m, "GetSnapshotStack", _param0)
	ret0, _ := ret[0].(interfaces.SnapshotStack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshotStack indicates an expected call of GetSnapshotStack
func (_mr *MockBackendClientMockRecorder) GetSnapshotStack(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSnapshotStack", reflect.TypeOf((*MockBackendClient)(nil).GetSnapshotStack), arg0)
}

// GetStack mocks base method
func (_m *MockBackendClient) GetStack(_param0 string) (types0.Stack, error) {
	ret := _m.ctrl.Call(_m, "GetStack", _param0)
	ret0, _ := ret[0].(types0.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStack indicates an expected call of GetStack
func (_mr *MockBackendClientMockRecorder) GetStack(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetStack", reflect.TypeOf((*MockBackendClient)(nil).GetStack), arg0)
}

// GetTask mocks base method
func (_m *MockBackendClient) GetTask(_param0 string) (swarm.Task, error) {
	ret := _m.ctrl.Call(_m, "GetTask", _param0)
	ret0, _ := ret[0].(swarm.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask
func (_mr *MockBackendClientMockRecorder) GetTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetTask", reflect.TypeOf((*MockBackendClient)(nil).GetTask), arg0)
}

// GetTasks mocks base method
func (_m *MockBackendClient) GetTasks(_param0 types.TaskListOptions) ([]swarm.Task, error) {
	ret := _m.ctrl.Call(_m, "GetTasks", _param0)
	ret0, _ := ret[0].([]swarm.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks
func (_mr *MockBackendClientMockRecorder) GetTasks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetTasks", reflect.TypeOf((*MockBackendClient)(nil).GetTasks), arg0)
}

// Info mocks base method
func (_m *MockBackendClient) Info() swarm.Info {
	ret := _m.ctrl.Call(_m, "Info")
	ret0, _ := ret[0].(swarm.Info)
	return ret0
}

// Info indicates an expected call of Info
func (_mr *MockBackendClientMockRecorder) Info() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Info", reflect.TypeOf((*MockBackendClient)(nil).Info))
}

// ListStacks mocks base method
func (_m *MockBackendClient) ListStacks() ([]types0.Stack, error) {
	ret := _m.ctrl.Call(_m, "ListStacks")
	ret0, _ := ret[0].([]types0.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStacks indicates an expected call of ListStacks
func (_mr *MockBackendClientMockRecorder) ListStacks() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ListStacks", reflect.TypeOf((*MockBackendClient)(nil).ListStacks))
}

// RemoveConfig mocks base method
func (_m *MockBackendClient) RemoveConfig(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveConfig", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveConfig indicates an expected call of RemoveConfig
func (_mr *MockBackendClientMockRecorder) RemoveConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveConfig", reflect.TypeOf((*MockBackendClient)(nil).RemoveConfig), arg0)
}

// RemoveNetwork mocks base method
func (_m *MockBackendClient) RemoveNetwork(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveNetwork", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNetwork indicates an expected call of RemoveNetwork
func (_mr *MockBackendClientMockRecorder) RemoveNetwork(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveNetwork", reflect.TypeOf((*MockBackendClient)(nil).RemoveNetwork), arg0)
}

// RemoveSecret mocks base method
func (_m *MockBackendClient) RemoveSecret(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveSecret", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecret indicates an expected call of RemoveSecret
func (_mr *MockBackendClientMockRecorder) RemoveSecret(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveSecret", reflect.TypeOf((*MockBackendClient)(nil).RemoveSecret), arg0)
}

// RemoveService mocks base method
func (_m *MockBackendClient) RemoveService(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveService", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveService indicates an expected call of RemoveService
func (_mr *MockBackendClientMockRecorder) RemoveService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveService", reflect.TypeOf((*MockBackendClient)(nil).RemoveService), arg0)
}

// SubscribeToEvents mocks base method
func (_m *MockBackendClient) SubscribeToEvents(_param0 time.Time, _param1 time.Time, _param2 filters.Args) ([]events.Message, chan interface{}) {
	ret := _m.ctrl.Call(_m, "SubscribeToEvents", _param0, _param1, _param2)
	ret0, _ := ret[0].([]events.Message)
	ret1, _ := ret[1].(chan interface{})
	return ret0, ret1
}

// SubscribeToEvents indicates an expected call of SubscribeToEvents
func (_mr *MockBackendClientMockRecorder) SubscribeToEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SubscribeToEvents", reflect.TypeOf((*MockBackendClient)(nil).SubscribeToEvents), arg0, arg1, arg2)
}

// UnsubscribeFromEvents mocks base method
func (_m *MockBackendClient) UnsubscribeFromEvents(_param0 chan interface{}) {
	_m.ctrl.Call(_m, "UnsubscribeFromEvents", _param0)
}

// UnsubscribeFromEvents indicates an expected call of UnsubscribeFromEvents
func (_mr *MockBackendClientMockRecorder) UnsubscribeFromEvents(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UnsubscribeFromEvents", reflect.TypeOf((*MockBackendClient)(nil).UnsubscribeFromEvents), arg0)
}

// UpdateConfig mocks base method
func (_m *MockBackendClient) UpdateConfig(_param0 string, _param1 uint64, _param2 swarm.ConfigSpec) error {
	ret := _m.ctrl.Call(_m, "UpdateConfig", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfig indicates an expected call of UpdateConfig
func (_mr *MockBackendClientMockRecorder) UpdateConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateConfig", reflect.TypeOf((*MockBackendClient)(nil).UpdateConfig), arg0, arg1, arg2)
}

// UpdateSecret mocks base method
func (_m *MockBackendClient) UpdateSecret(_param0 string, _param1 uint64, _param2 swarm.SecretSpec) error {
	ret := _m.ctrl.Call(_m, "UpdateSecret", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret
func (_mr *MockBackendClientMockRecorder) UpdateSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateSecret", reflect.TypeOf((*MockBackendClient)(nil).UpdateSecret), arg0, arg1, arg2)
}

// UpdateService mocks base method
func (_m *MockBackendClient) UpdateService(_param0 string, _param1 uint64, _param2 swarm.ServiceSpec, _param3 types.ServiceUpdateOptions, _param4 bool) (*types.ServiceUpdateResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateService", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(*types.ServiceUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService
func (_mr *MockBackendClientMockRecorder) UpdateService(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateService", reflect.TypeOf((*MockBackendClient)(nil).UpdateService), arg0, arg1, arg2, arg3, arg4)
}

// UpdateSnapshotStack mocks base method
func (_m *MockBackendClient) UpdateSnapshotStack(_param0 string, _param1 interfaces.SnapshotStack, _param2 uint64) (interfaces.SnapshotStack, error) {
	ret := _m.ctrl.Call(_m, "UpdateSnapshotStack", _param0, _param1, _param2)
	ret0, _ := ret[0].(interfaces.SnapshotStack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSnapshotStack indicates an expected call of UpdateSnapshotStack
func (_mr *MockBackendClientMockRecorder) UpdateSnapshotStack(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateSnapshotStack", reflect.TypeOf((*MockBackendClient)(nil).UpdateSnapshotStack), arg0, arg1, arg2)
}

// UpdateStack mocks base method
func (_m *MockBackendClient) UpdateStack(_param0 string, _param1 types0.StackSpec, _param2 uint64) error {
	ret := _m.ctrl.Call(_m, "UpdateStack", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStack indicates an expected call of UpdateStack
func (_mr *MockBackendClientMockRecorder) UpdateStack(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateStack", reflect.TypeOf((*MockBackendClient)(nil).UpdateStack), arg0, arg1, arg2)
}
