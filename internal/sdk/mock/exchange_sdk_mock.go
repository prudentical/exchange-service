// Code generated by MockGen. DO NOT EDIT.
// Source: internal/sdk/exchange_sdk.go
//
// Generated by this command:
//
//	mockgen -source=internal/sdk/exchange_sdk.go -destination=internal/sdk/mock/exchange_sdk_mock.go
//

// Package mock_sdk is a generated GoMock package.
package mock_sdk

import (
	model "exchange-service/internal/model"
	sdk "exchange-service/internal/sdk"
	reflect "reflect"
	time "time"

	decimal "github.com/shopspring/decimal"
	gomock "go.uber.org/mock/gomock"
)

// MockExchangeAPIClient is a mock of ExchangeAPIClient interface.
type MockExchangeAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeAPIClientMockRecorder
}

// MockExchangeAPIClientMockRecorder is the mock recorder for MockExchangeAPIClient.
type MockExchangeAPIClientMockRecorder struct {
	mock *MockExchangeAPIClient
}

// NewMockExchangeAPIClient creates a new mock instance.
func NewMockExchangeAPIClient(ctrl *gomock.Controller) *MockExchangeAPIClient {
	mock := &MockExchangeAPIClient{ctrl: ctrl}
	mock.recorder = &MockExchangeAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeAPIClient) EXPECT() *MockExchangeAPIClientMockRecorder {
	return m.recorder
}

// Currencies mocks base method.
func (m *MockExchangeAPIClient) Currencies() ([]model.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Currencies")
	ret0, _ := ret[0].([]model.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Currencies indicates an expected call of Currencies.
func (mr *MockExchangeAPIClientMockRecorder) Currencies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Currencies", reflect.TypeOf((*MockExchangeAPIClient)(nil).Currencies))
}

// GetExchange mocks base method.
func (m *MockExchangeAPIClient) GetExchange() model.Exchange {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchange")
	ret0, _ := ret[0].(model.Exchange)
	return ret0
}

// GetExchange indicates an expected call of GetExchange.
func (mr *MockExchangeAPIClientMockRecorder) GetExchange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchange", reflect.TypeOf((*MockExchangeAPIClient)(nil).GetExchange))
}

// HistoricPrice mocks base method.
func (m *MockExchangeAPIClient) HistoricPrice(pair model.Pair, time time.Time) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HistoricPrice", pair, time)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HistoricPrice indicates an expected call of HistoricPrice.
func (mr *MockExchangeAPIClientMockRecorder) HistoricPrice(pair, time any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HistoricPrice", reflect.TypeOf((*MockExchangeAPIClient)(nil).HistoricPrice), pair, time)
}

// Pairs mocks base method.
func (m *MockExchangeAPIClient) Pairs() ([]model.Pair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pairs")
	ret0, _ := ret[0].([]model.Pair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pairs indicates an expected call of Pairs.
func (mr *MockExchangeAPIClientMockRecorder) Pairs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pairs", reflect.TypeOf((*MockExchangeAPIClient)(nil).Pairs))
}

// PriceFor mocks base method.
func (m *MockExchangeAPIClient) PriceFor(pair model.Pair, amount, funds *decimal.Decimal, tradeType sdk.TradeType) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PriceFor", pair, amount, funds, tradeType)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PriceFor indicates an expected call of PriceFor.
func (mr *MockExchangeAPIClientMockRecorder) PriceFor(pair, amount, funds, tradeType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PriceFor", reflect.TypeOf((*MockExchangeAPIClient)(nil).PriceFor), pair, amount, funds, tradeType)
}

// MockExchangeAPIClientFactory is a mock of ExchangeAPIClientFactory interface.
type MockExchangeAPIClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeAPIClientFactoryMockRecorder
}

// MockExchangeAPIClientFactoryMockRecorder is the mock recorder for MockExchangeAPIClientFactory.
type MockExchangeAPIClientFactoryMockRecorder struct {
	mock *MockExchangeAPIClientFactory
}

// NewMockExchangeAPIClientFactory creates a new mock instance.
func NewMockExchangeAPIClientFactory(ctrl *gomock.Controller) *MockExchangeAPIClientFactory {
	mock := &MockExchangeAPIClientFactory{ctrl: ctrl}
	mock.recorder = &MockExchangeAPIClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeAPIClientFactory) EXPECT() *MockExchangeAPIClientFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockExchangeAPIClientFactory) Create(exchange model.Exchange) (sdk.ExchangeAPIClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", exchange)
	ret0, _ := ret[0].(sdk.ExchangeAPIClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockExchangeAPIClientFactoryMockRecorder) Create(exchange any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockExchangeAPIClientFactory)(nil).Create), exchange)
}
