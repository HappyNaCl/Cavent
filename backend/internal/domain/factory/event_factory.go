package factory

import (
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type EventFactory struct{}

func NewEventFactory() *EventFactory {
	return &EventFactory{}
}

func (ef *EventFactory) GetEvent(id uuid.UUID, createdById, title, eventType,
	ticketType, location, bannerUrl string, startTime int64,
	) *model.Event {
	return &model.Event{
		Id:          id,
		Title:       title,
		CreatedById: createdById,
		EventType:   eventType,
		TicketType:  ticketType,
		StartTime:   time.Unix(startTime, 0),
		Location:    location,
		BannerUrl:   bannerUrl,	
	}
}