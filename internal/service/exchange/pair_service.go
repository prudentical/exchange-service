package exchange

import (
	"errors"
	"exchange-service/internal/configuration"
	"exchange-service/internal/model"
	"exchange-service/internal/persistence"
	"exchange-service/internal/service"
	"exchange-service/internal/service/currency"
	"log/slog"
)

type PairService interface {
	GetAll(exchangeId int64, page int, size int) (persistence.Page[model.Pair], error)
	GetById(exchangeId int64, id int64) (model.Pair, error)
	Create(exchangeId int64, pair model.Pair) (model.Pair, error)
	Update(exchangeId int64, id int64, pair model.Pair) (model.Pair, error)
	Delete(exchangeId int64, id int64) error
	Merge(paris []model.Pair) error
}

type pairServiceImpl struct {
	dao      persistence.PairDAO
	exchange ExchangeService
	currency currency.CurrencyService
	config   configuration.Config
	logger   *slog.Logger
}

func NewPairService(dao persistence.PairDAO, exchange ExchangeService,
	currency currency.CurrencyService, config configuration.Config, logger *slog.Logger) PairService {
	return pairServiceImpl{dao, exchange, currency, config, logger}
}

func (s pairServiceImpl) GetAll(exchangeId int64, page int, size int) (persistence.Page[model.Pair], error) {
	s.logger.Debug("Get all")
	if page <= 0 {
		return persistence.Page[model.Pair]{}, service.InvalidPageError{}
	}
	maxSize := s.config.App.Pagination.MaxSize
	if size <= 0 || size > maxSize {
		return persistence.Page[model.Pair]{}, service.OutOfBoundPageSizeError{Max: maxSize}
	}
	_, err := s.exchange.GetById(exchangeId)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return persistence.Page[model.Pair]{}, service.NotFoundError{Type: model.Exchange{}, Id: exchangeId}
		}
		return persistence.Page[model.Pair]{}, err
	}
	pairs, err := s.dao.GetByExchangeId(exchangeId, page, size)
	if err != nil {
		return pairs, err
	}
	return pairs, err
}

func (s pairServiceImpl) GetById(exchangeId int64, id int64) (model.Pair, error) {
	_, err := s.exchange.GetById(exchangeId)
	if err != nil {
		return model.Pair{}, err
	}
	pair, err := s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return model.Pair{}, service.NotFoundError{Type: model.Pair{}, Id: id}
		}
		return model.Pair{}, err
	}
	if pair.ExchangeID != exchangeId {
		return model.Pair{}, service.ExchangePairMismatchError{}
	}
	return pair, err
}

func (s pairServiceImpl) Create(exchangeId int64, pair model.Pair) (model.Pair, error) {
	pair.ID = 0
	exchange, err := s.exchange.GetById(exchangeId)
	if err != nil {
		return model.Pair{}, err
	}
	pair.ExchangeID = exchange.ID
	pair.Exchange = exchange
	return s.dao.Create(pair)
}

func (s pairServiceImpl) Update(exchangeId int64, id int64, pair model.Pair) (model.Pair, error) {
	pair.ID = id
	exchange, err := s.exchange.GetById(exchangeId)
	if err != nil {
		return model.Pair{}, err
	}
	existing, err := s.dao.Get(id)
	if err != nil {
		return model.Pair{}, err
	}
	pair.ExchangeID = exchange.ID
	pair.Exchange = exchange
	pair.CreatedAt = existing.CreatedAt
	return s.dao.Update(pair)
}

func (s pairServiceImpl) Delete(exchangeId int64, id int64) error {
	_, err := s.exchange.GetById(exchangeId)
	if err != nil {
		return service.NotFoundError{Type: model.Exchange{}, Id: id}
	}
	_, err = s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return service.NotFoundError{Type: model.Pair{}, Id: id}
		}
		return err
	}
	return s.dao.Delete(id)
}

func (s pairServiceImpl) Merge(paris []model.Pair) error {
	s.logger.Debug("Merging pairs")
	for _, pair := range paris {
		s.logger.Debug("Checking", "pair", pair.Symbol, "base", pair.Base.Symbol, "quote", pair.Quote.Symbol)
		existing, err := s.dao.FindBy("symbol", pair.Symbol)
		if len(existing) > 0 {
			return nil
		}
		s.logger.Debug("Checking for existing pair", "pair", pair.Symbol, "found", len(existing))
		if err != nil {
			return err
		}
		base, err := s.currency.FindBySymbol(pair.Base.Symbol)
		s.logger.Debug("Checking for base currency", "base", pair.Base.Symbol, "found", len(base))
		if err != nil {
			return err
		}
		quote, err := s.currency.FindBySymbol(pair.Quote.Symbol)
		s.logger.Debug("Checking for quote currency", "quote", pair.Quote.Symbol, "found", len(quote))
		if err != nil {
			return err
		}
		if len(base) != 0 {
			pair.BaseID = base[0].ID
			pair.Base = model.Currency{}
		} else {
			created, err := s.currency.Create(pair.Base)
			s.logger.Debug("Created base currency", "base", pair.Base.Symbol, "id", created.ID)
			if err != nil {
				return err
			}
			pair.BaseID = created.ID
		}
		if len(quote) != 0 {
			pair.QuoteID = quote[0].ID
			pair.Quote = model.Currency{}
		} else {
			created, err := s.currency.Create(pair.Quote)
			s.logger.Debug("Created quote currency", "quote", pair.Quote.Symbol, "id", created.ID)
			if err != nil {
				return err
			}
			pair.QuoteID = created.ID
		}
		s.dao.Create(pair)
	}
	return nil
}
