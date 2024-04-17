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
	DateTime time.Time       `json:"date_time"`
	PairId   int             `json:"pair_id"`
	Amount   decimal.Decimal `json:"amount"`
	Price    decimal.Decimal `json:"price"`
}

func (o OrderRequest) toOrder(orderType service.OrderType) service.Order {
	return service.Order{
		BotId:    o.BotId,
		Amount:   o.Amount,
		Price:    o.Price,
		Type:     orderType,
		DateTime: time.Now(),
	}
}

type ExchangeOrderService interface {
	Buy(exchange sdk.ExchangeSDK, request OrderRequest) error
	Sell(exchange sdk.ExchangeSDK, request OrderRequest) error
}

type exchangeOrderServiceImpl struct {
	dao    persistence.ExchangeDAO
	pairs  PairService
	orders service.OrderService
}

func NewExchangeOrderService(dao persistence.ExchangeDAO, pairs PairService, orders service.OrderService) ExchangeOrderService {
	return exchangeOrderServiceImpl{dao: dao, pairs: pairs, orders: orders}
}

func (s exchangeOrderServiceImpl) Buy(exchange sdk.ExchangeSDK, request OrderRequest) error {
	pair, err := s.pairs.GetById(exchange.GetExchange().ID, request.PairId)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return service.NotFoundError{Type: model.Pair{}, Id: request.PairId}
		}
		return err
	}
	if request.Virtual {
		price, err := exchange.HistoricPrice(pair, request.DateTime)
		if err != nil {
			return err
		}
		request.Price = price
	} else {
		// Todo: Implement
		return errors.New("not implemented")
	}
	return s.orders.SaveOrder(request.toOrder(service.BuyOrder))
}

func (s exchangeOrderServiceImpl) Sell(exchange sdk.ExchangeSDK, request OrderRequest) error {
	pair, err := s.pairs.GetById(exchange.GetExchange().ID, request.PairId)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return service.NotFoundError{Type: model.Pair{}, Id: request.PairId}
		}
		return err
	}
	if request.Virtual {
		price, err := exchange.HistoricPrice(pair, request.DateTime)
		if err != nil {
			return err
		}
		request.Price = price
	} else {
		// Todo: Implement
		return errors.New("not implemented")
	}

	err = s.orders.SaveOrder(request.toOrder(service.SellOrder))
	return err
}
