package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    Id              uuid.UUID `gorm:"type:uuid;primaryKey"`
    CampusId        *uuid.UUID `gorm:"type:uuid"`
    Campus          Campus
    Provider        string `json:"provider" gorm:"not null"`
	ProviderId 	 	string `json:"providerId" gorm:"not null;uniqueIndex"`
    Email           string `json:"email"  gorm:"not null;uniqueIndex"`
    Name            string `json:"name" gorm:"not null"`
    Password        string `json:"-" gorm:"not null"`
    AvatarUrl       string `json:"avatarUrl" gorm:"not null"`
    FirstTimeLogin  bool   `json:"firstTimeLogin" gorm:"default:true"`
    Description     *string `json:"description"`
    Role            string `gorm:"default:'user'"`
    PhoneNumber     *string `json:"phoneNumber"`
    Address         *string `json:"address"`
	CreatedAt       time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

    Interests       []Category     `gorm:"many2many:user_interests"`
    Favorites       []UserFavorite `gorm:"foreignKey:UserId"`
    Tickets         []Ticket       `gorm:"foreignKey:UserId"`
    EventViews      []EventView    `gorm:"foreignKey:UserId"`
}
