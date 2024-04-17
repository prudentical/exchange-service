// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/exchange/exchange_price_service.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/exchange/exchange_price_service.go -destination=internal/service/exchange/mock/exchange_price_service_mock.go
//

// Package mock_exchange is a generated GoMock package.
package mock_exchange

import (
	sdk "exchange-service/internal/sdk"
	exchange "exchange-service/internal/service/exchange"
	reflect "reflect"

	decimal "github.com/shopspring/decimal"
	gomock "go.uber.org/mock/gomock"
)

// MockExchangePriceService is a mock of ExchangePriceService interface.
type MockExchangePriceService struct {
	ctrl     *gomock.Controller
	recorder *MockExchangePriceServiceMockRecorder
}

// MockExchangePriceServiceMockRecorder is the mock recorder for MockExchangePriceService.
type MockExchangePriceServiceMockRecorder struct {
	mock *MockExchangePriceService
}

// NewMockExchangePriceService creates a new mock instance.
func NewMockExchangePriceService(ctrl *gomock.Controller) *MockExchangePriceService {
	mock := &MockExchangePriceService{ctrl: ctrl}
	mock.recorder = &MockExchangePriceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangePriceService) EXPECT() *MockExchangePriceServiceMockRecorder {
	return m.recorder
}

// HistoricPrice mocks base method.
func (m *MockExchangePriceService) HistoricPrice(exchange sdk.ExchangeSDK, request exchange.PriceCheckRequest) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HistoricPrice", exchange, request)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HistoricPrice indicates an expected call of HistoricPrice.
func (mr *MockExchangePriceServiceMockRecorder) HistoricPrice(exchange, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HistoricPrice", reflect.TypeOf((*MockExchangePriceService)(nil).HistoricPrice), exchange, request)
}

// PriceFor mocks base method.
func (m *MockExchangePriceService) PriceFor(exchange sdk.ExchangeSDK, request exchange.PriceCheckRequest) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PriceFor", exchange, request)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PriceFor indicates an expected call of PriceFor.
func (mr *MockExchangePriceServiceMockRecorder) PriceFor(exchange, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PriceFor", reflect.TypeOf((*MockExchangePriceService)(nil).PriceFor), exchange, request)
}
