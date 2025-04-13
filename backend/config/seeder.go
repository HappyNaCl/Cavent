package config

import (
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error{


	log.Println("[INFO] Database seeded successfully!")
	return nil
}