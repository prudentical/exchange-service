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
	ApiKey    string           `json:"api_key"`
	DateTime  *time.Time       `json:"date_time"`
	Amount    *decimal.Decimal `json:"amount" validate:"required"`
	TradeType *sdk.TradeType   `json:"trade_type" validate:"required"`
}

type PriceCheckResponse struct {
	Price decimal.Decimal `json:"price"`
}

type PriceService interface {
	GetPrice(exchange sdk.ExchangeSDK, pairId int64, request PriceCheckRequest) (decimal.Decimal, error)
}
type priceServiceImpl struct {
	pairs PairService
}

func NewPriceService(pairs PairService) PriceService {
	return priceServiceImpl{pairs: pairs}
}

func (s priceServiceImpl) GetPrice(exchange sdk.ExchangeSDK, pairId int64, request PriceCheckRequest) (decimal.Decimal, error) {
	pair, err := s.pairs.GetById(exchange.GetExchange().ID, pairId)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return decimal.Decimal{}, service.NotFoundError{Type: model.Pair{}, Id: pairId}
		}
		return decimal.Decimal{}, err
	}
	if request.DateTime == nil {
		return exchange.PriceFor(pair, *request.Amount, *request.TradeType)
	}
	return exchange.HistoricPrice(pair, *request.DateTime)
}
