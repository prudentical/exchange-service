package exchange

type ExchangeService interface {
	ExchangeManageService
	ExchangeOrderService
	ExchangePriceService
}

type exchangeServiceImpl struct {
	ExchangeManageService
	ExchangeOrderService
	ExchangePriceService
}

func NewExchangeService(manager ExchangeManageService, orders ExchangeOrderService, prices ExchangePriceService) ExchangeService {
	return exchangeServiceImpl{manager, orders, prices}
}
