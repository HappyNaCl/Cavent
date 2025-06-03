package common

import "github.com/google/uuid"

type CampusResult struct {
	Id  		uuid.UUID `json:"id"`
	Name 		string    `json:"name"`
	LogoUrl 	string 	  `json:"logoUrl"`
	Description string 	  `json:"description"`
	InviteCode 	string	  `json:"inviteCode"`
}