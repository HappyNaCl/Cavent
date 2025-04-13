package config

import (
	"log"

	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error{
	db.Migrator().DropTable(&model.User{})
	db.Migrator().DropTable(&model.Tag{})
	db.Migrator().DropTable("user_interests")
	db.Migrator().DropTable(&model.TagType{})

	err := db.AutoMigrate(
		&model.User{},
		&model.Tag{},
		&model.TagType{},
	)

	if err != nil {
		return err
	}
	
	log.Println("[INFO] Database successfully migrated!")
	return nil
}