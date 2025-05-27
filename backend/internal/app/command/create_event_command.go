package command

import (
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
)

type CreateEventCommand struct {
	CreatedById string
	Title   	string
	EventType 	string
	TicketType 	string
	StartTime 	time.Time
	EndTime   	*time.Time
	Location  	string
	Description *string
	BannerBytes []byte
	BannerExt  	string
	Ticket     []common.TicketResult
} 

type CreateEventCommandResult struct {
	Result *common.EventResult
}