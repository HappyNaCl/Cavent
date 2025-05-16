package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
    Id          uuid.UUID `gorm:"type:uuid;primaryKey"`
    Title       string
    CampusId    uuid.UUID
    Campus      Campus
    CategoryId  uuid.UUID
    Category    Category
    EventType   string
    TicketType  string
    StartTime   time.Time
    EndTime     time.Time
    Location    string
    Description string
    BannerUrl   string
    CreatedAt   time.Time
    UpdatedAt   time.Time
	
    TicketTypes []TicketType `gorm:"foreignKey:EventId"`
    Favorites   []UserFavorite `gorm:"foreignKey:EventId"`
    Views       []EventView    `gorm:"foreignKey:EventId"`
}