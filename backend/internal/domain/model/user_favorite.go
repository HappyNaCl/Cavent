package model

import (
	"time"

	"github.com/google/uuid"
)

type UserFavorite struct {
	UserId  	string 		`gorm:"primaryKey"`
	EventId 	uuid.UUID 	`gorm:"type:uuid;primaryKey"`

	Event    	Event     	`gorm:"foreignKey:EventId"`
	User     	User      	`gorm:"foreignKey:UserId"`

	CreatedAt 	time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt 	time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}