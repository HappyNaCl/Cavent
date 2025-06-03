package command

import "github.com/google/uuid"

type CreateEventViewCommand struct {
	EventID uuid.UUID `json:"eventId"`
	UserID  *string `json:"userId"`
}

type CreateEventViewCommandResult struct {}