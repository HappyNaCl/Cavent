package postgres

import (
	"log"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    err := db.Migrator().DropTable(
        &model.UserFavorite{},
        &model.EventView{},
        &model.Ticket{},
        &model.TicketType{},
        &model.Event{},
        &model.Category{},
        &model.CategoryType{},
        &model.User{},
        &model.Campus{},
        "user_interests",
    )

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
        &model.Campus{},
        &model.User{},
        &model.CategoryType{},
        &model.Category{},
        &model.UserFavorite{},
        &model.Event{},
        &model.TicketType{},
        &model.Ticket{},
        &model.EventView{},
    )
	
    if err != nil {
        log.Fatal("failed to migrate:", err)
    }
}