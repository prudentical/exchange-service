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

type ExchangeAPIClient interface {
	Currencies() ([]model.Currency, error)
	Pairs() ([]model.Pair, error)
	PriceFor(pair model.Pair, amount *decimal.Decimal, funds *decimal.Decimal, tradeType TradeType) (decimal.Decimal, error)
	HistoricPrice(pair model.Pair, time time.Time) (decimal.Decimal, error)
	GetExchange() model.Exchange
}

type ExchangeAPIClientFactory interface {
	Create(exchange model.Exchange) (ExchangeAPIClient, error)
}

type exchangeAPIClientFactoryImpl struct {
	wallexClient WallexClient
}

func NewExchangeAPIClientFactory(wallexClient WallexClient) ExchangeAPIClientFactory {
	return exchangeAPIClientFactoryImpl{wallexClient}
}

func (f exchangeAPIClientFactoryImpl) Create(exchange model.Exchange) (ExchangeAPIClient, error) {
	switch ExchangeType(exchange.Name) {
	case Wallex:
		return newWallexAPIClient(exchange, f.wallexClient), nil
	default:
		return wallexAPI{}, NoImplementationFoundError{exchange.Name}
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
