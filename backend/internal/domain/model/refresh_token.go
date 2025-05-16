package model

import "time"

type RefreshToken struct {
	UserId string `json:"userId" gorm:"type:uuid;primaryKey"`
	Token  string `json:"token" gorm:"type:varchar(100);uniqueIndex"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	User      *User      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}