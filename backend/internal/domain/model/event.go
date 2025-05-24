package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
    Id          uuid.UUID `gorm:"type:uuid;primaryKey"`
    Title       string    `json:"title" gorm:"not null"`
    CampusId    uuid.UUID `gorm:"type:uuid"`
    Campus      Campus
    CategoryId  uuid.UUID `gorm:"type:uuid"`
    Category    Category
    EventType   string    `json:"eventType" gorm:"not null"`
    TicketType  string    `json:"ticketType" gorm:"not null"`
    StartTime   time.Time
    EndTime     time.Time
    Langitude   float32
    Longitude   float32     
    Address     string
    Description string
    BannerUrl   string
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	
    TicketTypes []TicketType `gorm:"foreignKey:EventId"`
    Favorites   []UserFavorite `gorm:"foreignKey:EventId"`
    Views       []EventView    `gorm:"foreignKey:EventId"`
}