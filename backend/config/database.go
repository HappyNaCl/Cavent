package config

import (
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error){
	url := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        zap.L().Sugar().Errorf("Error connecting to db: %v", err)
		return nil, err
    }

	return db, nil
}