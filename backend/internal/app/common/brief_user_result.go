package common

import "github.com/google/uuid"

type BriefUserResult struct {
	Id           	uuid.UUID `json:"id"`
	CampusId   		*uuid.UUID `json:"campusId,omitempty"`
	Provider    	string `json:"provider"`
	Name		 	string `json:"name"`
	Email        	string `json:"email"`
	AvatarUrl   	string `json:"avatarUrl"`
}