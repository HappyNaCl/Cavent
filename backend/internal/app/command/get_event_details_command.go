package command

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/google/uuid"
)

type GetEventDetailsCommand struct {
	EventID uuid.UUID `json:"eventId"`
	UserId  *string  `json:"userId,omitempty"`
}

type GetEventDetailsCommandResult struct {
	Result *common.EventResult
}