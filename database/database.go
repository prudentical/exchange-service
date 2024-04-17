package database

import (
	"exchange-service/configuration"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbCon *gorm.DB

func connect(config configuration.Config, logger *slog.Logger) *gorm.DB {
	logger.Info("Connecting to the database")
	database := config.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		database.Host, database.User, database.Password, database.Name, database.Port, database.SSL, database.Timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewDatabaseConnection(config configuration.Config, logger *slog.Logger) *gorm.DB {
	if dbCon == nil {
		dbCon = setupDB(config, logger)
	}
	return dbCon
}

func setupDB(config configuration.Config, logger *slog.Logger) *gorm.DB {
	conn := connect(config, logger)
	db, err := conn.DB()
	if err != nil {
		panic(err)
	}
	err = migrateSchema(conn, logger)
	if err != nil {
		panic(err)
	}
	err = insertSeed(conn, logger)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(config.Database.Connection.Idle)
	db.SetMaxOpenConns(config.Database.Connection.Open)
	db.SetConnMaxLifetime(time.Hour)
	if config.Logging.Level == configuration.Debug {
		conn = conn.Debug()
	}
	return conn
}

func migrateSchema(db *gorm.DB, logger *slog.Logger) error {
	logger.Debug("Migrating database schema")
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	driver, err := pg.WithInstance(sqlDB, &pg.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != migrate.ErrNoChange {
		return err
	}
	return nil
}
