package common

import "github.com/google/uuid"

type BriefCampusResult struct {
	Id          uuid.UUID `json:"id"`
	Name        string `json:"name"`
	ProfileUrl 	string `json:"profileUrl"`
}