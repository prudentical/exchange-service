package exchange

import (
	"errors"
	"exchange-service/internal/model"
	"exchange-service/internal/persistence"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service"
)

type ExchangeService interface {
	GetAll() ([]model.Exchange, error)
	GetAllWithPage(page int, size int) (persistence.Page[model.Exchange], error)
	GetById(id int64) (model.Exchange, error)
	Update(id int64, exchange model.Exchange) (model.Exchange, error)
}

type exchangeServiceImpl struct {
	dao     persistence.ExchangeDAO
	factory sdk.ExchangeSDKFactory
}

func NewExchangeService(dao persistence.ExchangeDAO, factory sdk.ExchangeSDKFactory) ExchangeService {
	return exchangeServiceImpl{dao, factory}
}

func (s exchangeServiceImpl) GetAll() ([]model.Exchange, error) {
	return s.dao.GetAll()
}

func (s exchangeServiceImpl) GetAllWithPage(page int, size int) (persistence.Page[model.Exchange], error) {
	return s.dao.GetAllWithPage(page, size)
}

func (s exchangeServiceImpl) GetById(id int64) (model.Exchange, error) {
	exchange, err := s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return model.Exchange{}, service.NotFoundError{Type: model.Exchange{}, Id: id}
		}
		return model.Exchange{}, err
	}
	return exchange, nil
}

func (s exchangeServiceImpl) Update(id int64, exchange model.Exchange) (model.Exchange, error) {
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
