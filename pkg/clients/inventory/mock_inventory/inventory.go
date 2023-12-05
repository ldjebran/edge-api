// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/clients/inventory/client.go

// Package mock_inventory is a generated GoMock package.
package mock_inventory

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	inventory "github.com/redhatinsights/edge-api/pkg/clients/inventory"
)

// MockClientInterface is a mock of ClientInterface interface.
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface.
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance.
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// BuildURL mocks base method.
func (m *MockClientInterface) BuildURL(parameters *inventory.Params) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildURL", parameters)
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildURL indicates an expected call of BuildURL.
func (mr *MockClientInterfaceMockRecorder) BuildURL(parameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildURL", reflect.TypeOf((*MockClientInterface)(nil).BuildURL), parameters)
}

// ReturnDeviceListByID mocks base method.
func (m *MockClientInterface) ReturnDeviceListByID(deviceIDs []string) (inventory.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnDeviceListByID", deviceIDs)
	ret0, _ := ret[0].(inventory.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReturnDeviceListByID indicates an expected call of ReturnDeviceListByID.
func (mr *MockClientInterfaceMockRecorder) ReturnDeviceListByID(deviceIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnDeviceListByID", reflect.TypeOf((*MockClientInterface)(nil).ReturnDeviceListByID), deviceIDs)
}

// ReturnDevices mocks base method.
func (m *MockClientInterface) ReturnDevices(parameters *inventory.Params) (inventory.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnDevices", parameters)
	ret0, _ := ret[0].(inventory.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReturnDevices indicates an expected call of ReturnDevices.
func (mr *MockClientInterfaceMockRecorder) ReturnDevices(parameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnDevices", reflect.TypeOf((*MockClientInterface)(nil).ReturnDevices), parameters)
}

// ReturnDevicesByID mocks base method.
func (m *MockClientInterface) ReturnDevicesByID(deviceID string) (inventory.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnDevicesByID", deviceID)
	ret0, _ := ret[0].(inventory.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReturnDevicesByID indicates an expected call of ReturnDevicesByID.
func (mr *MockClientInterfaceMockRecorder) ReturnDevicesByID(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnDevicesByID", reflect.TypeOf((*MockClientInterface)(nil).ReturnDevicesByID), deviceID)
}

// ReturnDevicesByTag mocks base method.
func (m *MockClientInterface) ReturnDevicesByTag(tag string) (inventory.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnDevicesByTag", tag)
	ret0, _ := ret[0].(inventory.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReturnDevicesByTag indicates an expected call of ReturnDevicesByTag.
func (mr *MockClientInterfaceMockRecorder) ReturnDevicesByTag(tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnDevicesByTag", reflect.TypeOf((*MockClientInterface)(nil).ReturnDevicesByTag), tag)
}
