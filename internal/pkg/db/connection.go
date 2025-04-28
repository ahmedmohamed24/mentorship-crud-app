package db

import (
	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

func NewDBClient(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
