package tasks

import (
	"encoding/json"

	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/hibiken/asynq"
)

type TicketCheckoutPayload struct {
	UserId   string  `json:"userId"`
	Tickets  []*common.CheckoutTicketResult `json:"tickets"`
}

func NewTicketCheckoutPayload(userId string, tickets []*common.CheckoutTicketResult) (*asynq.Task, error){
	payload := &TicketCheckoutPayload{
		UserId:  userId,
		Tickets: tickets,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	
	return asynq.NewTask(TypeTicketCheckout, data), nil
}