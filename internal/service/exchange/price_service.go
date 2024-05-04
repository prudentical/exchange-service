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
	ApiKey    string           `json:"apiKey" query:"apiKey"`
	DateTime  *time.Time       `json:"datetime" query:"datetime"`
	Amount    *decimal.Decimal `json:"amount" query:"amount"`
	Funds     *decimal.Decimal `json:"funds" query:"funds"`
	TradeType sdk.TradeType    `json:"tradeType" query:"tradeType" validate:"required"`
}

type PriceCheckResponse struct {
	Price decimal.Decimal `json:"price"`
}

type PriceService interface {
	GetPrice(exchange sdk.ExchangeAPIClient, pairId int64, request PriceCheckRequest) (decimal.Decimal, error)
}
type priceServiceImpl struct {
	pairs PairService
}

func NewPriceService(pairs PairService) PriceService {
	return priceServiceImpl{pairs: pairs}
}

func (s priceServiceImpl) GetPrice(exchange sdk.ExchangeAPIClient, pairId int64, request PriceCheckRequest) (decimal.Decimal, error) {
	pair, err := s.pairs.GetById(exchange.GetExchange().ID, pairId)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return decimal.Decimal{}, service.NotFoundError{Type: model.Pair{}, Id: pairId}
		}
		return decimal.Decimal{}, err
	}
	if request.DateTime == nil {
		return exchange.PriceFor(pair, request.Amount, request.Funds, request.TradeType)
	}
	return exchange.HistoricPrice(pair, *request.DateTime)
}
