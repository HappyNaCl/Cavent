package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
    Id              uuid.UUID `gorm:"type:uuid;primaryKey"`
    CategoryTypeId  uuid.UUID
    CategoryType    CategoryType
    Name            string    `json:"name" gorm:"unique;not null"`
	CreatedAt       time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

    Users      []User     `gorm:"many2many:user_interests"`
    Events     []Event    `gorm:"foreignKey:CategoryId"`
}