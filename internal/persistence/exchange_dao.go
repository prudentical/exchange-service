package persistence

import (
	"errors"
	"exchange-service/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type ExchangeDAO interface {
	GetAll() ([]model.Exchange, error)
	GetAllWithPage(page int, size int) (Page[model.Exchange], error)
	Create(exchange model.Exchange) (model.Exchange, error)
	Update(exchange model.Exchange) (model.Exchange, error)
	Get(id int64) (model.Exchange, error)
	Delete(id int64) error
	FindBy(field string, value any) ([]model.Exchange, error)
}

type exchangeDAOImpl struct {
	db *gorm.DB
}

func NewExchangeDAO(conn *gorm.DB) ExchangeDAO {
	return exchangeDAOImpl{conn}
}

func (dao exchangeDAOImpl) GetAllWithPage(page int, size int) (Page[model.Exchange], error) {
	var exchanges []model.Exchange
	tx := dao.db.Scopes(Paginate(page, size)).Find(&exchanges)
	if tx.Error != nil {
		return Page[model.Exchange]{}, tx.Error
	}
	var total int64
	dao.db.Model(model.Exchange{}).Count(&total)
	return Page[model.Exchange]{
		List:  exchanges,
		Page:  page,
		Size:  len(exchanges),
		Total: total,
	}, nil
}

func (dao exchangeDAOImpl) GetAll() ([]model.Exchange, error) {
	var exchanges []model.Exchange
	tx := dao.db.Find(&exchanges)
	if tx.Error != nil {
		return []model.Exchange{}, tx.Error
	}
	return exchanges, nil
}

func (dao exchangeDAOImpl) Get(id int64) (model.Exchange, error) {
	var exchange model.Exchange
	tx := dao.db.First(&exchange, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.Exchange{}, RecordNotFoundError{}
		}
		return model.Exchange{}, tx.Error
	}
	return exchange, nil
}
func (dao exchangeDAOImpl) Create(exchange model.Exchange) (model.Exchange, error) {
	tx := dao.db.Create(&exchange)
	if tx.Error != nil {
		return model.Exchange{}, tx.Error
	}
	return exchange, nil
}

func (dao exchangeDAOImpl) Update(exchange model.Exchange) (model.Exchange, error) {
	tx := dao.db.Save(&exchange)
	if tx.Error != nil {
		return model.Exchange{}, tx.Error
	}
	return exchange, nil
}

func (dao exchangeDAOImpl) Delete(id int64) error {
	tx := dao.db.Delete(model.Exchange{}, id)
	return tx.Error
}

func (dao exchangeDAOImpl) FindBy(field string, value any) ([]model.Exchange, error) {
	var exchanges []model.Exchange
	tx := dao.db.Find(&exchanges, fmt.Sprintf("%s = ?", field), value)
	if tx.Error != nil {
		return []model.Exchange{}, tx.Error
	}
	return exchanges, nil

}
