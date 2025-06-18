package handler

import (
	"context"
	"encoding/json"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/tasks"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)


type TicketCheckoutHandler struct {
	ticketService *service.TicketService
}

func NewTicketCheckoutHandler(db *gorm.DB) *TicketCheckoutHandler {
	return &TicketCheckoutHandler{
		ticketService: service.NewTicketService(db),
	}
}

func (h *TicketCheckoutHandler) Handle(ctx context.Context, task *asynq.Task) error {
	zap.L().Sugar().Infof("Processing task: %s", task.Type())

	var payload tasks.TicketCheckoutPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	for _, ticket := range payload.Tickets {
		com := &command.CreateTicketCommand{
			TicketId: ticket.Id,
			UserId:  payload.UserId,
		}

		for i := 0; i < ticket.Quantity; i++ {
			if _, err := h.ticketService.CreateTicket(ctx, com); err != nil {
				zap.L().Sugar().Errorf("Failed to create ticket: %v", err)
				return err
			}
		}
	}

	return nil
}