package common

import (
	"time"

	"github.com/google/uuid"
)


type UserResult struct {
	Id          	string
    CampusId        *uuid.UUID
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
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

func (u UserResult) ToBrief() BriefUserResult {
    return BriefUserResult{
        Id:        u.Id,
        CampusId:  u.CampusId,
        Provider:  u.Provider,
        Name:      u.Name,
        Email:     u.Email,
        AvatarUrl: u.AvatarUrl,
        FirstTimeLogin: u.FirstTimeLogin,
        Role:      u.Role,
    }
}