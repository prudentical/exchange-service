// Code generated by MockGen. DO NOT EDIT.
// Source: internal/api/pair_handler.go
//
// Generated by this command:
//
//	mockgen -source=internal/api/pair_handler.go -destination=internal/api/mock/pair_handler_mock.go
//

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockPairHandler is a mock of PairHandler interface.
type MockPairHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPairHandlerMockRecorder
}

// MockPairHandlerMockRecorder is the mock recorder for MockPairHandler.
type MockPairHandlerMockRecorder struct {
	mock *MockPairHandler
}

// NewMockPairHandler creates a new mock instance.
func NewMockPairHandler(ctrl *gomock.Controller) *MockPairHandler {
	mock := &MockPairHandler{ctrl: ctrl}
	mock.recorder = &MockPairHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPairHandler) EXPECT() *MockPairHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPairHandler) Create(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPairHandlerMockRecorder) Create(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPairHandler)(nil).Create), c)
}

// Delete mocks base method.
func (m *MockPairHandler) Delete(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPairHandlerMockRecorder) Delete(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPairHandler)(nil).Delete), c)
}

// GetAll mocks base method.
func (m *MockPairHandler) GetAll(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPairHandlerMockRecorder) GetAll(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPairHandler)(nil).GetAll), c)
}

// GetById mocks base method.
func (m *MockPairHandler) GetById(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetById indicates an expected call of GetById.
func (mr *MockPairHandlerMockRecorder) GetById(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPairHandler)(nil).GetById), c)
}

// HandleRoutes mocks base method.
func (m *MockPairHandler) HandleRoutes(e *echo.Echo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleRoutes", e)
}

// HandleRoutes indicates an expected call of HandleRoutes.
func (mr *MockPairHandlerMockRecorder) HandleRoutes(e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRoutes", reflect.TypeOf((*MockPairHandler)(nil).HandleRoutes), e)
}

// Update mocks base method.
func (m *MockPairHandler) Update(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPairHandlerMockRecorder) Update(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPairHandler)(nil).Update), c)
}
