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

type OrderRequest struct {
	BotId    int             `json:"bot_id"`
	ApiKey   string          `json:"api_key"`
	Virtual  bool            `json:"virtual"`
	DateTime *time.Time      `json:"date_time"`
	Amount   decimal.Decimal `json:"amount"`
	Price    decimal.Decimal `json:"price"`
	Type     sdk.TradeType   `json:"type"`
}

func (o OrderRequest) toOrder() service.Order {
	return service.Order{
		BotId:    o.BotId,
		Amount:   o.Amount,
		Price:    o.Price,
		Type:     o.Type,
		DateTime: time.Now(),
	}
}

type OrderService interface {
	Order(exchange sdk.ExchangeSDK, pairId int64, request OrderRequest) error
}

type orderServiceImpl struct {
	dao    persistence.ExchangeDAO
	pairs  PairService
	orders service.OrderService
}

func NewOrderService(dao persistence.ExchangeDAO, pairs PairService, orders service.OrderService) OrderService {
	return orderServiceImpl{dao: dao, pairs: pairs, orders: orders}
}

func (s orderServiceImpl) Order(exchange sdk.ExchangeSDK, pairId int64, request OrderRequest) error {
	pair, err := s.pairs.GetById(exchange.GetExchange().ID, pairId)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return service.NotFoundError{Type: model.Pair{}, Id: pairId}
		}
		return err
	}

	if request.Virtual {
		var price decimal.Decimal
		var err error
		if request.DateTime == nil {
			price, err = exchange.PriceFor(pair, request.Amount, request.Type)
		} else {
			price, err = exchange.HistoricPrice(pair, *request.DateTime)
		}
		if err != nil {
			return err
		}
		request.Price = price
	} else {
		switch request.Type {
		case sdk.Buy:
			// Todo: Implement buy
		case sdk.Sell:
			// Todo: Implement sell
		default:
			// TODO: return custom error
		}
		return errors.New("not implemented")
	}
	return s.orders.SaveOrder(request.toOrder())
}
