package persistence

import (
	"errors"
	"exchange-service/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type CurrencyDAO interface {
	GetAll(page int, size int) (Page[model.Currency], error)
	Create(currency model.Currency) (model.Currency, error)
	Update(currency model.Currency) (model.Currency, error)
	FindBy(field string, value any) ([]model.Currency, error)
	Get(id int) (model.Currency, error)
	Delete(id int) error
}

type currencyDAOImpl struct {
	db *gorm.DB
}

func NewCurrencyDAO(conn *gorm.DB) CurrencyDAO {
	return currencyDAOImpl{conn}
}

func (dao currencyDAOImpl) GetAll(page int, size int) (Page[model.Currency], error) {
	var currencies []model.Currency
	tx := dao.db.Scopes(Paginate(page, size)).Find(&currencies)
	if tx.Error != nil {
		return Page[model.Currency]{}, tx.Error
	}
	var total int64
	dao.db.Model(model.Currency{}).Count(&total)
	return Page[model.Currency]{
		List:  currencies,
		Page:  page,
		Size:  len(currencies),
		Total: total,
	}, nil
}

func (dao currencyDAOImpl) Get(id int) (model.Currency, error) {
	var currency model.Currency
	tx := dao.db.First(&currency, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.Currency{}, RecordNotFoundError{}
		}
		return model.Currency{}, tx.Error
	}
	return currency, nil
}

func (dao currencyDAOImpl) Create(currency model.Currency) (model.Currency, error) {
	tx := dao.db.Create(&currency)
	if tx.Error != nil {
		return model.Currency{}, tx.Error
	}
	return currency, nil
}

func (dao currencyDAOImpl) Update(currency model.Currency) (model.Currency, error) {
	tx := dao.db.Save(&currency)
	if tx.Error != nil {
		return model.Currency{}, tx.Error
	}
	return currency, nil
}

func (dao currencyDAOImpl) Delete(id int) error {
	tx := dao.db.Delete(model.Currency{}, id)
	return tx.Error
}

func (dao currencyDAOImpl) FindBy(field string, value any) ([]model.Currency, error) {
	var currency []model.Currency
	tx := dao.db.Find(&currency, fmt.Sprintf("%s = ?", field), value)
	if tx.Error != nil {
		return []model.Currency{}, tx.Error
	}
	return currency, nil

}
