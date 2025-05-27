package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
    Id          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    Title       string    `json:"title" gorm:"not null"`
    CampusId    uuid.UUID `json:"campusId" gorm:"type:uuid"`
    CreatedById string  `json:"createdById" gorm:"not null"`
    CreatedBy   User    `json:"createdBy" gorm:"foreignKey:CreatedById"`
    Campus      Campus     
    EventType   string    `json:"eventType" gorm:"not null"`
    TicketType  string    `json:"ticketType" gorm:"not null"`
    StartTime   time.Time
    EndTime     *time.Time
    Location    string    
    Description *string
    BannerUrl   string
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	
    TicketTypes []TicketType    `gorm:"foreignKey:EventId"`
    Favorites   []UserFavorite  `gorm:"foreignKey:EventId"`
    Views       []EventView     `gorm:"foreignKey:EventId"`
    Categories  []Category      `gorm:"many2many:event_categories;"`
}