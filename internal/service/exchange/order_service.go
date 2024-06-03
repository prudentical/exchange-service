package exchange

import (
	"errors"
	"exchange-service/internal/dto"
	"exchange-service/internal/persistence"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service"

	"github.com/shopspring/decimal"
)

type OrderService interface {
	Order(exchange sdk.ExchangeAPIClient, pairId int64, request dto.OrderDTO) error
}

type orderServiceImpl struct {
	dao    persistence.ExchangeDAO
	orders service.OrderService
	pairs  PairService
}

func NewOrderService(dao persistence.ExchangeDAO, orders service.OrderService, pairs PairService) OrderService {
	return orderServiceImpl{dao: dao, orders: orders, pairs: pairs}
}

func (s orderServiceImpl) Order(exchange sdk.ExchangeAPIClient, pairId int64, request dto.OrderDTO) error {
	if request.Amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return service.OrderAmountRequiredError{}
	}
	pair, err := s.pairs.GetById(exchange.GetExchange().ID, pairId)
	if err != nil {
		return err
	}
	if request.IsVirtual {
		if request.DateTime != nil {
			price, err := exchange.HistoricPrice(pair, *request.DateTime)
			if err != nil {
				return err
			}
			request.Price = price
		} else {
			price, err := exchange.PriceFor(pair, &request.Amount, nil, request.Type)
			if err != nil {
				return err
			}
			request.Price = price
		}
		request.FilledAmount = request.Amount
		request.Status = dto.Fulfilled
		request.InternalId = "VIRT"
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
	return s.orders.SaveOrder(request)
}
