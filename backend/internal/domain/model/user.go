package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    Id              string `gorm:"primaryKey"`
    CampusId        *uuid.UUID `gorm:"type:uuid"`
    Campus          Campus
    Provider        string `json:"provider" gorm:"not null"`
    Email           string `json:"email"  gorm:"not null;uniqueIndex"`
    Name            string `json:"name" gorm:"not null"`
    Password        string `json:"-"`
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
