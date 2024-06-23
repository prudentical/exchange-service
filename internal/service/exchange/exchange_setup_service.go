package exchange

import (
	"exchange-service/internal/sdk"
	"exchange-service/internal/service/currency"
)

type ExchangeSetupService interface {
	Setup() error
	addCurrencies(exchangeSDK sdk.ExchangeAPIClient) error
	addPairs(exchangeSDK sdk.ExchangeAPIClient) error
}
type exchangeSetupServiceImpl struct {
	factory    sdk.ExchangeAPIClientFactory
	currencies currency.CurrencyService
	exchanges  ExchangeService
	pairs      PairService
}

func NewExchangeSetupService(factory sdk.ExchangeAPIClientFactory, currencies currency.CurrencyService, exchanges ExchangeService, pairs PairService) ExchangeSetupService {
	return exchangeSetupServiceImpl{factory, currencies, exchanges, pairs}
}

func (s exchangeSetupServiceImpl) Setup() error {
	exchanges, err := s.exchanges.GetAll()
	if err != nil {
		return err
	}
	for _, exchange := range exchanges {
		exchangeSDK, err := s.factory.Create(exchange)
		if err != nil {
			return err
		}
		err = s.addCurrencies(exchangeSDK)
		if err != nil {
			return err
		}
		err = s.addPairs(exchangeSDK)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s exchangeSetupServiceImpl) addCurrencies(exchangeSDK sdk.ExchangeAPIClient) error {
	currencies, err := exchangeSDK.Currencies()
	if err != nil {
		return err
	}
	return s.currencies.Merge(currencies)
}

func (s exchangeSetupServiceImpl) addPairs(exchangeSDK sdk.ExchangeAPIClient) error {
	pairs, err := exchangeSDK.Pairs()
	if err != nil {
		return err
	}
	return s.pairs.Merge(pairs)
}
