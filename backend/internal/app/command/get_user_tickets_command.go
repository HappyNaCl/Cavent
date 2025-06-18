package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetUserTicketsCommand struct {
	UserId string `json:"userId" validate:"required"`
}

type GetUserTicketsCommandResponse struct {
	Tickets []*common.UserTicketResult `json:"tickets"`
}