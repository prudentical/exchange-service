package exchange

import (
	"errors"
	"exchange-service/internal/model"
	"exchange-service/internal/persistence"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service"
)

type ExchangeManageService interface {
	GetAll() ([]model.Exchange, error)
	GetAllWithPage(page int, size int) (persistence.Page[model.Exchange], error)
	GetById(id int) (model.Exchange, error)
	Create(exchange model.Exchange) (model.Exchange, error)
	Update(id int, exchange model.Exchange) (model.Exchange, error)
	Delete(id int) error
}

type exchangeManagerServiceImpl struct {
	dao     persistence.ExchangeDAO
	factory sdk.ExchangeSDKFactory
}

func NewExchangeManagerService(dao persistence.ExchangeDAO, factory sdk.ExchangeSDKFactory) ExchangeManageService {
	return exchangeManagerServiceImpl{dao, factory}
}

func (s exchangeManagerServiceImpl) GetAllWithPage(page int, size int) (persistence.Page[model.Exchange], error) {
	return s.dao.GetAllWithPage(page, size)
}

func (s exchangeManagerServiceImpl) GetAll() ([]model.Exchange, error) {
	return s.dao.GetAll()
}

func (s exchangeManagerServiceImpl) GetById(id int) (model.Exchange, error) {
	exchange, err := s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return model.Exchange{}, service.NotFoundError{Type: model.Exchange{}, Id: id}
		}
		return model.Exchange{}, err
	}
	return exchange, nil
}

func (s exchangeManagerServiceImpl) Create(exchange model.Exchange) (model.Exchange, error) {
	_, err := s.factory.Create(exchange)
	if err != nil {
		return model.Exchange{}, err
	}
	exchange.ID = 0
	return s.dao.Create(exchange)
}

func (s exchangeManagerServiceImpl) Update(id int, exchange model.Exchange) (model.Exchange, error) {
	_, err := s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return model.Exchange{}, service.NotFoundError{Type: model.Exchange{}, Id: id}
		}
		return model.Exchange{}, err
	}
	exchange.ID = id
	return s.dao.Update(exchange)
}

func (s exchangeManagerServiceImpl) Delete(id int) error {
	_, err := s.dao.Get(id)
	if err != nil && errors.Is(err, persistence.RecordNotFoundError{}) {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return service.NotFoundError{Type: model.Exchange{}, Id: id}
		}
		return err
	}
	return s.dao.Delete(id)
}
