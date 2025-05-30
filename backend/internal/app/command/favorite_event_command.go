package command

import "github.com/google/uuid"

type FavoriteEventCommand struct {
	EventId uuid.UUID `json:"eventId" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
}

type FavoriteEventCommandResult struct {
	Result int64 `json:"eventCount"`
}