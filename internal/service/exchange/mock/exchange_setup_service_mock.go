// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/exchange/exchange_setup_service.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/exchange/exchange_setup_service.go -destination=internal/service/exchange/mock/exchange_setup_service_mock.go
//

// Package mock_exchange is a generated GoMock package.
package mock_exchange

import (
	sdk "exchange-service/internal/sdk"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockExchangeSetupService is a mock of ExchangeSetupService interface.
type MockExchangeSetupService struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeSetupServiceMockRecorder
}

// MockExchangeSetupServiceMockRecorder is the mock recorder for MockExchangeSetupService.
type MockExchangeSetupServiceMockRecorder struct {
	mock *MockExchangeSetupService
}

// NewMockExchangeSetupService creates a new mock instance.
func NewMockExchangeSetupService(ctrl *gomock.Controller) *MockExchangeSetupService {
	mock := &MockExchangeSetupService{ctrl: ctrl}
	mock.recorder = &MockExchangeSetupServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeSetupService) EXPECT() *MockExchangeSetupServiceMockRecorder {
	return m.recorder
}

// Setup mocks base method.
func (m *MockExchangeSetupService) Setup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup.
func (mr *MockExchangeSetupServiceMockRecorder) Setup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockExchangeSetupService)(nil).Setup))
}

// addCurrencies mocks base method.
func (m *MockExchangeSetupService) addCurrencies(exchangeSDK sdk.ExchangeAPIClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "addCurrencies", exchangeSDK)
	ret0, _ := ret[0].(error)
	return ret0
}

// addCurrencies indicates an expected call of addCurrencies.
func (mr *MockExchangeSetupServiceMockRecorder) addCurrencies(exchangeSDK any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addCurrencies", reflect.TypeOf((*MockExchangeSetupService)(nil).addCurrencies), exchangeSDK)
}

// addPairs mocks base method.
func (m *MockExchangeSetupService) addPairs(exchangeSDK sdk.ExchangeAPIClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "addPairs", exchangeSDK)
	ret0, _ := ret[0].(error)
	return ret0
}

// addPairs indicates an expected call of addPairs.
func (mr *MockExchangeSetupServiceMockRecorder) addPairs(exchangeSDK any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addPairs", reflect.TypeOf((*MockExchangeSetupService)(nil).addPairs), exchangeSDK)
}
