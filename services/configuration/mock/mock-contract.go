// Code generated by MockGen. DO NOT EDIT.
// Source: services/configuration/contract.go

// Package mock_configuration is a generated GoMock package.
package mock_configuration

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConfigurationContract is a mock of ConfigurationContract interface.
type MockConfigurationContract struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationContractMockRecorder
}

// MockConfigurationContractMockRecorder is the mock recorder for MockConfigurationContract.
type MockConfigurationContractMockRecorder struct {
	mock *MockConfigurationContract
}

// NewMockConfigurationContract creates a new mock instance.
func NewMockConfigurationContract(ctrl *gomock.Controller) *MockConfigurationContract {
	mock := &MockConfigurationContract{ctrl: ctrl}
	mock.recorder = &MockConfigurationContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationContract) EXPECT() *MockConfigurationContractMockRecorder {
	return m.recorder
}

// GetDatabaseCollectionName mocks base method.
func (m *MockConfigurationContract) GetDatabaseCollectionName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabaseCollectionName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatabaseCollectionName indicates an expected call of GetDatabaseCollectionName.
func (mr *MockConfigurationContractMockRecorder) GetDatabaseCollectionName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabaseCollectionName", reflect.TypeOf((*MockConfigurationContract)(nil).GetDatabaseCollectionName))
}

// GetDatabaseConnectionString mocks base method.
func (m *MockConfigurationContract) GetDatabaseConnectionString() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabaseConnectionString")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatabaseConnectionString indicates an expected call of GetDatabaseConnectionString.
func (mr *MockConfigurationContractMockRecorder) GetDatabaseConnectionString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabaseConnectionString", reflect.TypeOf((*MockConfigurationContract)(nil).GetDatabaseConnectionString))
}

// GetDatabaseName mocks base method.
func (m *MockConfigurationContract) GetDatabaseName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabaseName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatabaseName indicates an expected call of GetDatabaseName.
func (mr *MockConfigurationContractMockRecorder) GetDatabaseName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabaseName", reflect.TypeOf((*MockConfigurationContract)(nil).GetDatabaseName))
}

// GetGrpcHost mocks base method.
func (m *MockConfigurationContract) GetGrpcHost() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrpcHost")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGrpcHost indicates an expected call of GetGrpcHost.
func (mr *MockConfigurationContractMockRecorder) GetGrpcHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrpcHost", reflect.TypeOf((*MockConfigurationContract)(nil).GetGrpcHost))
}

// GetGrpcPort mocks base method.
func (m *MockConfigurationContract) GetGrpcPort() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrpcPort")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGrpcPort indicates an expected call of GetGrpcPort.
func (mr *MockConfigurationContractMockRecorder) GetGrpcPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrpcPort", reflect.TypeOf((*MockConfigurationContract)(nil).GetGrpcPort))
}

// GetHttpHost mocks base method.
func (m *MockConfigurationContract) GetHttpHost() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHttpHost")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHttpHost indicates an expected call of GetHttpHost.
func (mr *MockConfigurationContractMockRecorder) GetHttpHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttpHost", reflect.TypeOf((*MockConfigurationContract)(nil).GetHttpHost))
}

// GetHttpPort mocks base method.
func (m *MockConfigurationContract) GetHttpPort() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHttpPort")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHttpPort indicates an expected call of GetHttpPort.
func (mr *MockConfigurationContractMockRecorder) GetHttpPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttpPort", reflect.TypeOf((*MockConfigurationContract)(nil).GetHttpPort))
}

// GetJwksURL mocks base method.
func (m *MockConfigurationContract) GetJwksURL() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJwksURL")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJwksURL indicates an expected call of GetJwksURL.
func (mr *MockConfigurationContractMockRecorder) GetJwksURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJwksURL", reflect.TypeOf((*MockConfigurationContract)(nil).GetJwksURL))
}

// GetK3SDockerImage mocks base method.
func (m *MockConfigurationContract) GetK3SDockerImage() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetK3SDockerImage")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetK3SDockerImage indicates an expected call of GetK3SDockerImage.
func (mr *MockConfigurationContractMockRecorder) GetK3SDockerImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetK3SDockerImage", reflect.TypeOf((*MockConfigurationContract)(nil).GetK3SDockerImage))
}
