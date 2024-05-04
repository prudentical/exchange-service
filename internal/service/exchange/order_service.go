package exchange

import (
	"errors"
	"exchange-service/internal/persistence"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service"
	"time"

	"github.com/shopspring/decimal"
)

type OrderRequest struct {
	BotId    int             `json:"botId"`
	ApiKey   string          `json:"apiKey"`
	Virtual  bool            `json:"virtual"`
	DateTime *time.Time      `json:"datetime"`
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
	Order(exchange sdk.ExchangeAPIClient, pairId int64, request OrderRequest) error
}

type orderServiceImpl struct {
	dao    persistence.ExchangeDAO
	orders service.OrderService
}

func NewOrderService(dao persistence.ExchangeDAO, orders service.OrderService) OrderService {
	return orderServiceImpl{dao: dao, orders: orders}
}

func (s orderServiceImpl) Order(exchange sdk.ExchangeAPIClient, pairId int64, request OrderRequest) error {
	if !request.Virtual {
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
