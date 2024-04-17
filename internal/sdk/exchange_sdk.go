package sdk

import (
	"exchange-service/internal/model"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type TradeType string

const (
	Sell TradeType = "Sell"
	Buy  TradeType = "Buy"
)

type InvalidTradeType struct {
	tradeType string
}

func (e InvalidTradeType) Error() string {
	return fmt.Sprintf("Invalid trade type {%s}", e.tradeType)
}

type ExchangeSDK interface {
	Currencies() ([]model.Currency, error)
	Pairs() ([]model.Pair, error)
	PriceFor(pair model.Pair, amount decimal.Decimal, tradeType TradeType) (decimal.Decimal, error)
	HistoricPrice(pair model.Pair, time time.Time) (decimal.Decimal, error)
	GetExchange() model.Exchange
}

type ExchangeSDKFactory interface {
	Create(exchange model.Exchange) (ExchangeSDK, error)
}

type exchangeSDKFactoryImpl struct {
}

func NewExchangeSDKFactory() ExchangeSDKFactory {
	return exchangeSDKFactoryImpl{}
}

func (exchangeSDKFactoryImpl) Create(exchange model.Exchange) (ExchangeSDK, error) {
	switch ExchangeType(exchange.Name) {
	case Wallex:
		return newWallexSDK(exchange), nil
	default:
		return wallexSDK{}, NoImplementationFoundError{exchange.Name}
	}
}

type ExchangeType string

const (
	Wallex ExchangeType = "Wallex"
)

type NoImplementationFoundError struct {
	Exchange string
}

func (e NoImplementationFoundError) Error() string {
	return fmt.Sprintf("Invalid exchange {%s}", e.Exchange)
}
