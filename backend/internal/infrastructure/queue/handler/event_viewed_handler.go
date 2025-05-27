package handler

import (
	"context"
	"encoding/json"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/tasks"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type EventViewedHandler struct {
	analyticService *service.AnalyticsService
}

func NewEventViewedHandler(db *gorm.DB) *EventViewedHandler {
	return &EventViewedHandler{
		analyticService: service.NewAnalyticsService(db),
	}
}

func (h *EventViewedHandler) Handle(ctx context.Context, task *asynq.Task) error {
	var payload tasks.EventViewPayload

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	com := &command.CreateEventViewCommand{
		EventID: payload.EventId,
		UserID:  payload.UserId,
	}

	_, err := h.analyticService.CreateEventView(com)
	if err != nil {
		return err
	}

	return nil
}