// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
)

// MockManagerInterface is a mock of ManagerInterface interface.
type MockManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockManagerInterfaceMockRecorder
}

// MockManagerInterfaceMockRecorder is the mock recorder for MockManagerInterface.
type MockManagerInterfaceMockRecorder struct {
	mock *MockManagerInterface
}

// NewMockManagerInterface creates a new mock instance.
func NewMockManagerInterface(ctrl *gomock.Controller) *MockManagerInterface {
	mock := &MockManagerInterface{ctrl: ctrl}
	mock.recorder = &MockManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagerInterface) EXPECT() *MockManagerInterfaceMockRecorder {
	return m.recorder
}

// ClearPCIAddressFolder mocks base method.
func (m *MockManagerInterface) ClearPCIAddressFolder() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearPCIAddressFolder")
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearPCIAddressFolder indicates an expected call of ClearPCIAddressFolder.
func (mr *MockManagerInterfaceMockRecorder) ClearPCIAddressFolder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearPCIAddressFolder", reflect.TypeOf((*MockManagerInterface)(nil).ClearPCIAddressFolder))
}

// GetCheckPointNodeState mocks base method.
func (m *MockManagerInterface) GetCheckPointNodeState() (*v1.SriovNetworkNodeState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckPointNodeState")
	ret0, _ := ret[0].(*v1.SriovNetworkNodeState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckPointNodeState indicates an expected call of GetCheckPointNodeState.
func (mr *MockManagerInterfaceMockRecorder) GetCheckPointNodeState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckPointNodeState", reflect.TypeOf((*MockManagerInterface)(nil).GetCheckPointNodeState))
}

// LoadPfsStatus mocks base method.
func (m *MockManagerInterface) LoadPfsStatus(pciAddress string) (*v1.Interface, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPfsStatus", pciAddress)
	ret0, _ := ret[0].(*v1.Interface)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadPfsStatus indicates an expected call of LoadPfsStatus.
func (mr *MockManagerInterfaceMockRecorder) LoadPfsStatus(pciAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPfsStatus", reflect.TypeOf((*MockManagerInterface)(nil).LoadPfsStatus), pciAddress)
}

// SaveLastPfAppliedStatus mocks base method.
func (m *MockManagerInterface) SaveLastPfAppliedStatus(PfInfo *v1.Interface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLastPfAppliedStatus", PfInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLastPfAppliedStatus indicates an expected call of SaveLastPfAppliedStatus.
func (mr *MockManagerInterfaceMockRecorder) SaveLastPfAppliedStatus(PfInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLastPfAppliedStatus", reflect.TypeOf((*MockManagerInterface)(nil).SaveLastPfAppliedStatus), PfInfo)
}

// WriteCheckpointFile mocks base method.
func (m *MockManagerInterface) WriteCheckpointFile(arg0 *v1.SriovNetworkNodeState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteCheckpointFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteCheckpointFile indicates an expected call of WriteCheckpointFile.
func (mr *MockManagerInterfaceMockRecorder) WriteCheckpointFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteCheckpointFile", reflect.TypeOf((*MockManagerInterface)(nil).WriteCheckpointFile), arg0)
}
