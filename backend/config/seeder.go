package config

import (
	"log"

	"github.com/HappyNaCl/Cavent/backend/config/seeder"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error{
	tagSeeder := seeder.NewTagSeeder()
	if err := tagSeeder.Seed(db); err != nil {
		log.Printf("[ERROR] Failed to seed database: %v", err)
		panic(err)
	}

	log.Println("[INFO] Database seeded successfully!")
	return nil
}