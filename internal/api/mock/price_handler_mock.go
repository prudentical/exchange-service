// Code generated by MockGen. DO NOT EDIT.
// Source: internal/api/price_handler.go
//
// Generated by this command:
//
//	mockgen -source=internal/api/price_handler.go -destination=internal/api/mock/price_handler_mock.go
//

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockPriceHandler is a mock of PriceHandler interface.
type MockPriceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPriceHandlerMockRecorder
}

// MockPriceHandlerMockRecorder is the mock recorder for MockPriceHandler.
type MockPriceHandlerMockRecorder struct {
	mock *MockPriceHandler
}

// NewMockPriceHandler creates a new mock instance.
func NewMockPriceHandler(ctrl *gomock.Controller) *MockPriceHandler {
	mock := &MockPriceHandler{ctrl: ctrl}
	mock.recorder = &MockPriceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceHandler) EXPECT() *MockPriceHandlerMockRecorder {
	return m.recorder
}

// GetPrice mocks base method.
func (m *MockPriceHandler) GetPrice(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockPriceHandlerMockRecorder) GetPrice(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockPriceHandler)(nil).GetPrice), c)
}

// HandleRoutes mocks base method.
func (m *MockPriceHandler) HandleRoutes(e *echo.Echo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleRoutes", e)
}

// HandleRoutes indicates an expected call of HandleRoutes.
func (mr *MockPriceHandlerMockRecorder) HandleRoutes(e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRoutes", reflect.TypeOf((*MockPriceHandler)(nil).HandleRoutes), e)
}
