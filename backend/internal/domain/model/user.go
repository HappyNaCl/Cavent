package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    Id              uuid.UUID `gorm:"type:uuid;primaryKey"`
    CampusId        uuid.UUID
    Campus          Campus
    Provider        string
	ProviderId 	 	string `gorm:"uniqueIndex"`
    Email           string
    Name            string
    Password        string
    AvatarUrl       string
    FirstTimeLogin  bool
    Description     string
    Role            string
    PhoneNumber     string
    Address         string
    CreatedAt       time.Time
    UpdatedAt       time.Time

    Interests       []Category     `gorm:"many2many:user_interests"`
    Favorites       []UserFavorite `gorm:"foreignKey:UserId"`
    Tickets         []Ticket       `gorm:"foreignKey:UserId"`
    EventViews      []EventView    `gorm:"foreignKey:UserId"`

	RefreshToken    RefreshToken   `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}
