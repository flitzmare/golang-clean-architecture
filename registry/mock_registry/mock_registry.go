// Code generated by MockGen. DO NOT EDIT.
// Source: registry/registry.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	gomock "github.com/golang/mock/gomock"
	controller "golang-clean-architecture/interface/controller"
	reflect "reflect"
)

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// NewAppController mocks base method.
func (m *MockRegistry) NewAppController() controller.AppController {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAppController")
	ret0, _ := ret[0].(controller.AppController)
	return ret0
}

// NewAppController indicates an expected call of NewAppController.
func (mr *MockRegistryMockRecorder) NewAppController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAppController", reflect.TypeOf((*MockRegistry)(nil).NewAppController))
}
