package common

import "github.com/google/uuid"

type BriefUserResult struct {
	Id           	string `json:"id"`
	CampusId   		*uuid.UUID `json:"campusId,omitempty"`
	Provider    	string `json:"provider"`
	Name		 	string `json:"name"`
	Email        	string `json:"email"`
	AvatarUrl   	string `json:"avatarUrl"`
	FirstTimeLogin 	bool   `json:"firstTimeLogin"`
	Role         	string `json:"role"`
}