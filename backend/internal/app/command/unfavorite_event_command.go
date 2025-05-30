package command

import "github.com/google/uuid"

type UnfavoriteEventCommand struct {
	EventId uuid.UUID `json:"eventId" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
}

type UnfavoriteEventCommandResult struct {
	Result int64 `json:"eventCount"`
}