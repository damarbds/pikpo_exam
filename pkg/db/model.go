package db

import "gorm.io/gorm"

type IDatabaseAdapter interface {
	Insert(model interface{}) (int, error)
	Update(model interface{}, id uint) (interface{}, error)
	ReadById(model interface{}, id uint) (interface{}, error)
	ReadAll(models interface{}) (interface{}, error)
}

type PostgresDbHandler struct {
	DB *gorm.DB
}
