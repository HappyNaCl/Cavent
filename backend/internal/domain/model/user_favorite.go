package model

import (
	"time"

	"github.com/google/uuid"
)

type UserFavorite struct {
	UserId  	uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	EventId 	uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	CreatedAt 	time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt 	time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}