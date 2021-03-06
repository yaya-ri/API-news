// Code generated by MockGen. DO NOT EDIT.
// Source: module/v1/controller/controller_contract.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockNewsControllerInterface is a mock of NewsControllerInterface interface
type MockNewsControllerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNewsControllerInterfaceMockRecorder
}

// MockNewsControllerInterfaceMockRecorder is the mock recorder for MockNewsControllerInterface
type MockNewsControllerInterfaceMockRecorder struct {
	mock *MockNewsControllerInterface
}

// NewMockNewsControllerInterface creates a new mock instance
func NewMockNewsControllerInterface(ctrl *gomock.Controller) *MockNewsControllerInterface {
	mock := &MockNewsControllerInterface{ctrl: ctrl}
	mock.recorder = &MockNewsControllerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNewsControllerInterface) EXPECT() *MockNewsControllerInterfaceMockRecorder {
	return m.recorder
}

// Store mocks base method
func (m *MockNewsControllerInterface) Store(context *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Store", context)
}

// Store indicates an expected call of Store
func (mr *MockNewsControllerInterfaceMockRecorder) Store(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockNewsControllerInterface)(nil).Store), context)
}

// Find mocks base method
func (m *MockNewsControllerInterface) Find(context *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Find", context)
}

// Find indicates an expected call of Find
func (mr *MockNewsControllerInterfaceMockRecorder) Find(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockNewsControllerInterface)(nil).Find), context)
}
