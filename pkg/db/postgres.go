package db

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgresDb(dsn string) (*gorm.DB, IDatabaseAdapter, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	handler := newPostgresDbHandler(db)
	return db, handler, nil
}

func newPostgresDbHandler(postgresDb *gorm.DB) IDatabaseAdapter {
	return &PostgresDbHandler{
		DB: postgresDb,
	}
}

func (handler *PostgresDbHandler) Insert(model interface{}) (int, error) {
	log.Debug().Msg("[PostgresDBHandler] query insert")
	err := handler.DB.Create(model).Error

	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (handler *PostgresDbHandler) Update(model interface{}, id uint) (interface{}, error) {
	log.Debug().Msg("[PostgresDBHandler] query update")
	err := handler.DB.Save(model).Error

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (handler *PostgresDbHandler) ReadById(model interface{}, id uint) (interface{}, error) {
	log.Debug().Msg("[PostgresDBHandler] query any")
	err := handler.DB.First(&model, id).Error

	if err != nil {
		return false, err
	}

	return model, nil
}

func (handler *PostgresDbHandler) ReadAll(models interface{}) (interface{}, error) {
	log.Debug().Msg("[PostgresDBHandler] query any")
	err := handler.DB.Find(&models).Error

	if err != nil {
		return nil, err
	}

	return models, nil
}
