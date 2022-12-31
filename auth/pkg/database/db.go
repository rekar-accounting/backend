package database

import (
	"github.com/rekar-accounting/backend/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	config.Load()
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {

		return nil, err
	}
	return db, nil

}
