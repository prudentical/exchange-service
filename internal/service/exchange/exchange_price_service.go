package exchange

import (
	"errors"
	"exchange-service/internal/model"
	"exchange-service/internal/persistence"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service"
	"time"

	"github.com/shopspring/decimal"
)

type PriceCheckRequest struct {
	ApiKey    string          `json:"api_key"`
	DateTime  time.Time       `json:"date_time"`
	PairId    int             `json:"pair_id" validate:"required"`
	Amount    decimal.Decimal `json:"amount" validate:"required"`
	TradeType sdk.TradeType   `json:"trade_type" validate:"required"`
}

type PriceCheckResponse struct {
	Price decimal.Decimal `json:"price"`
}

type ExchangePriceService interface {
	PriceFor(exchange sdk.ExchangeSDK, request PriceCheckRequest) (decimal.Decimal, error)
	HistoricPrice(exchange sdk.ExchangeSDK, request PriceCheckRequest) (decimal.Decimal, error)
}
type exchangePriceServiceImpl struct {
	pairs PairService
}

func NewExchangePriceService(pairs PairService) ExchangePriceService {
	return exchangePriceServiceImpl{pairs: pairs}
}

func (s exchangePriceServiceImpl) PriceFor(exchange sdk.ExchangeSDK, request PriceCheckRequest) (decimal.Decimal, error) {
	pair, err := s.pairs.GetById(exchange.GetExchange().ID, request.PairId)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return decimal.Decimal{}, service.NotFoundError{Type: model.Pair{}, Id: request.PairId}
		}
		return decimal.Decimal{}, err
	}
	price, err := exchange.PriceFor(pair, request.Amount, request.TradeType)
	return price, err
}

func (s exchangePriceServiceImpl) HistoricPrice(exchange sdk.ExchangeSDK, request PriceCheckRequest) (decimal.Decimal, error) {
	pair, err := s.pairs.GetById(exchange.GetExchange().ID, request.PairId)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return decimal.Decimal{}, service.NotFoundError{Type: model.Pair{}, Id: request.PairId}
		}
		return decimal.Decimal{}, err
	}
	price, err := exchange.HistoricPrice(pair, request.DateTime)
	return price, err
}
