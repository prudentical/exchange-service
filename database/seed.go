package database

import (
	"exchange-service/internal/model"
	"log/slog"

	"gorm.io/gorm"
)

var _exchanges = []model.Exchange{
	{Name: "Wallex"},
}

func insertSeed(conn *gorm.DB, logger *slog.Logger) error {
	logger.Debug("Inserting seed data")
	err := insertExchanges(conn, logger)
	if err != nil {
		return err
	}
	return nil
}

func insertExchanges(conn *gorm.DB, logger *slog.Logger) error {
	logger.Debug("Insert exchanges")
	var count int64
	conn.Model(model.Exchange{}).Count(&count)
	if count > 0 {
		return nil
	}
	for _, exchange := range _exchanges {
		trx := conn.Create(&exchange)
		if trx.Error != nil {
			return trx.Error
		}
	}
	return nil
}
