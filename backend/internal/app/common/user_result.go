package common

import (
	"time"

	"github.com/google/uuid"
)


type UserResult struct {
	Id          	uuid.UUID `gorm:"type:uuid;primaryKey"`
    CampusId        *uuid.UUID `gorm:"type:uuid"`
    Provider        string
	ProviderId 	 	string
    Email           string 
    Name            string 
    AvatarUrl       string 
    FirstTimeLogin  bool  
    Description     *string
    Role            string 
    PhoneNumber     *string 
    Address         *string 
	CreatedAt 		time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt 		time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u UserResult) ToBrief() BriefUserResult {
    return BriefUserResult{
        Id:        u.Id,
        CampusId:  u.CampusId,
        Provider:  u.Provider,
        Name:      u.Name,
        Email:     u.Email,
        AvatarUrl: u.AvatarUrl,
    }
}