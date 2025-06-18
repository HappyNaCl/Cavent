package common

import "github.com/google/uuid"

type UserTicketResult struct {
	EventId   uuid.UUID `json:"eventId"`
	EventTitle string `json:"eventTitle"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Tickets   []TicketResult `json:"tickets"`
}