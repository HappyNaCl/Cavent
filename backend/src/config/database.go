package config

import (
	"log"
	"os"

	"github.com/HappyNaCl/Cavent/src/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func ConnectDatabase() error{
	url := os.Getenv("DATABASE_URL")
	
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	
	if err != nil {
		return err
	}

	log.Println("[INFO] Database successfully connected!")
	Database = db


	db.AutoMigrate(&domain.User{})
	return nil
}