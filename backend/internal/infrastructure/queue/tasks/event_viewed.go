package tasks

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

type EventViewPayload struct {
	UserId    string `json:"userId"`
	EventId   uuid.UUID `json:"eventId"`
}

func NewEventViewPayload(userId string, eventId uuid.UUID) (*asynq.Task, error) {
	payload, err := json.Marshal(EventViewPayload{
		UserId:  userId,
		EventId: eventId,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeEventView, payload), nil
}