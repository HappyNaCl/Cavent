package model

type Tag struct {
	Id string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
	TagTypeId string `json:"tagTypeId" gorm:"not null"`
	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime"`
	Users []User `json:"users" gorm:"many2many:user_interests;joinForeignKey:TagId;joinReferences:UserId"`
}