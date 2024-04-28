package currency

import (
	"errors"
	"exchange-service/internal/configuration"
	"exchange-service/internal/model"
	"exchange-service/internal/persistence"
	"exchange-service/internal/service"
)

type CurrencyService interface {
	GetAll(page int, size int) (persistence.Page[model.Currency], error)
	GetById(id int) (model.Currency, error)
	Create(currency model.Currency) (model.Currency, error)
	Update(id int, currency model.Currency) (model.Currency, error)
	Delete(id int) error
	Merge(currencies []model.Currency) error
	FindBySymbol(symbol string) ([]model.Currency, error)
}

type currencyServiceImpl struct {
	dao    persistence.CurrencyDAO
	config configuration.Config
}

func NewCurrencyService(dao persistence.CurrencyDAO, config configuration.Config) CurrencyService {
	return currencyServiceImpl{dao, config}
}

func (s currencyServiceImpl) GetAll(page int, size int) (persistence.Page[model.Currency], error) {
	if page <= 0 || size <= 0 {
		return persistence.Page[model.Currency]{}, service.InvalidPageParameterError{}
	}
	if size > s.config.App.Pagination.MaxSize {
		return persistence.Page[model.Currency]{}, service.OutOfBoundPageSizeError{Max: s.config.App.Pagination.MaxSize}
	}
	return s.dao.GetAll(page, size)
}

func (s currencyServiceImpl) GetById(id int) (model.Currency, error) {
	currency, err := s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return model.Currency{}, service.NotFoundError{Type: model.Currency{}, Id: id}
		}
		return model.Currency{}, err
	}
	return currency, nil
}

func (s currencyServiceImpl) Create(currency model.Currency) (model.Currency, error) {
	currency.ID = 0
	return s.dao.Create(currency)
}

func (s currencyServiceImpl) Update(id int, currency model.Currency) (model.Currency, error) {
	_, err := s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return model.Currency{}, service.NotFoundError{Type: model.Currency{}, Id: id}
		}
		return model.Currency{}, err
	}
	currency.ID = id
	return s.dao.Update(currency)
}

func (s currencyServiceImpl) Delete(id int) error {
	_, err := s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return service.NotFoundError{Type: model.Currency{}, Id: id}
		}
		return err
	}
	return s.dao.Delete(id)
}

func (s currencyServiceImpl) FindBySymbol(symbol string) ([]model.Currency, error) {
	found, err := s.dao.FindBy("symbol", symbol)
	if err != nil {
		return []model.Currency{}, err
	}
	return found, nil
}

func (s currencyServiceImpl) Merge(currencies []model.Currency) error {
	for _, currency := range currencies {
		existing, err := s.FindBySymbol(currency.Symbol)
		if err != nil {
			return err
		}
		if len(existing) == 0 {
			s.dao.Create(currency)
		}
	}
	return nil
}
