package persistence

import (
	"errors"
	"exchange-service/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type PairDAO interface {
	GetByExchangeId(exchangeId int, page int, size int) (Page[model.Pair], error)
	Create(pair model.Pair) (model.Pair, error)
	Update(pair model.Pair) (model.Pair, error)
	Get(id int) (model.Pair, error)
	Delete(id int) error
	FindBy(field string, value any) ([]model.Pair, error)
}

type pairDAOImpl struct {
	db *gorm.DB
}

func NewPairDAO(conn *gorm.DB) PairDAO {
	return pairDAOImpl{conn}
}

func (dao pairDAOImpl) GetByExchangeId(exchangeId int, page int, size int) (Page[model.Pair], error) {
	var pairs []model.Pair
	// TODO: apply page limit
	tx := dao.db.Scopes(Paginate(page, size)).Find(&pairs, "exchange_id = ?", exchangeId)
	if tx.Error != nil {
		return Page[model.Pair]{}, tx.Error
	}
	var total int64
	dao.db.Model(model.Pair{}).Count(&total)
	return Page[model.Pair]{
		List:  pairs,
		Page:  page,
		Size:  len(pairs),
		Total: total,
	}, nil
}

func (dao pairDAOImpl) Get(id int) (model.Pair, error) {
	var pair model.Pair
	tx := dao.db.First(&pair, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.Pair{}, RecordNotFoundError{}
		}
		return model.Pair{}, tx.Error
	}
	return pair, nil
}

func (dao pairDAOImpl) Create(pair model.Pair) (model.Pair, error) {
	tx := dao.db.Create(&pair)
	if tx.Error != nil {
		return model.Pair{}, tx.Error
	}
	return pair, nil
}

func (dao pairDAOImpl) Update(pair model.Pair) (model.Pair, error) {
	tx := dao.db.Save(&pair)
	if tx.Error != nil {
		return model.Pair{}, tx.Error
	}
	return pair, nil
}

func (dao pairDAOImpl) Delete(id int) error {
	tx := dao.db.Delete(model.Pair{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dao pairDAOImpl) FindBy(field string, value any) ([]model.Pair, error) {
	var pairs []model.Pair
	tx := dao.db.Find(&pairs, fmt.Sprintf("%s = ?", field), value)
	if tx.Error != nil {
		return []model.Pair{}, tx.Error
	}
	return pairs, nil

}
