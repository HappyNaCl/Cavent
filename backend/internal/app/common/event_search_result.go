package common

import "github.com/google/uuid"

type EventSearchResult struct {
	Id 		uuid.UUID `json:"id"`
	Title 	string `json:"title"`
	StartTime int64 `json:"startTime"`
	Location string `json:"location"`
}