package model

import (
	"time"

	"github.com/google/uuid"
)

type CategoryType struct {
    Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
    Categories []Category `gorm:"foreignKey:TagTypeID"`
}