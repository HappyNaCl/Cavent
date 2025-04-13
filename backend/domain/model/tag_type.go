package model

type TagType struct {
	Id string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime"`
	Tags []Tag `json:"tags" gorm:"foreignKey:TagTypeId;references:Id"`
}