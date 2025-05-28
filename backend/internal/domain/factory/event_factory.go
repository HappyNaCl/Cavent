package factory

import (
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type EventFactory struct{}

func NewEventFactory() *EventFactory {
	return &EventFactory{}
}

func (ef *EventFactory) GetEvent(id uuid.UUID, createdById, title, eventType,
	ticketType, location, bannerUrl string, startTime int64, tickets []common.TicketResult,
	) *model.Event {
	ticketModels := make([]model.TicketType, 0, len(tickets))
	for _, ticket := range tickets {
		ticketModels = append(ticketModels, model.TicketType{
			Id:          uuid.New(),
			Name:        ticket.Name,
			Price: 	 	 ticket.Price,
			Quantity:    ticket.Quantity,
			EventId:     id,
		})
	}


	return &model.Event{
		Id:          id,
		Title:       title,
		CreatedById: createdById,
		EventType:   eventType,
		TicketType:  ticketType,
		StartTime:   time.Unix(startTime, 0),
		Location:    location,
		BannerUrl:   bannerUrl,
		TicketTypes: ticketModels,	
	}
}