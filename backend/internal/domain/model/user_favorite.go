package model

import (
	"time"

	"github.com/google/uuid"
)

type UserFavorite struct {
	UserId  	uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	EventId 	uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	createdAt 	time.Time 	`gorm:"autoCreateTime"`
	updatedAt 	time.Time 	`gorm:"autoUpdateTime"`
}