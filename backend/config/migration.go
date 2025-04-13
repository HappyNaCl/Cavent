package config

import (
	"log"

	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error{
	db.Migrator().DropTable(&model.User{})

	err := db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		return err
	}
	
	log.Println("[INFO] Database successfully migrated!")
	return nil
}