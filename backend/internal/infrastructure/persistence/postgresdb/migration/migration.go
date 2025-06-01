package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    err := db.Migrator().DropTable(
        &model.UserFavorite{},
        &model.EventView{},
        &model.Ticket{},
        &model.TicketType{},
        "event_categories",
        &model.Event{},
        &model.Category{},
        &model.CategoryType{},
        &model.User{},
        &model.Campus{},
        "user_interests",
    )

	if err != nil {
		return err
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
        return err
    }

    err = db.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm;").Error
    if err != nil {
        return err
    }

    err = db.Exec("CREATE INDEX idx_trgm_title ON events USING gin (title gin_trgm_ops);").Error
    if err != nil {
        return err
    }

    zap.L().Sugar().Info("Database migrated successfully!")
    return nil
}
