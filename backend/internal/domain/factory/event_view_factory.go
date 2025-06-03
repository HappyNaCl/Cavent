package factory

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type EventViewFactory struct{}

func NewEventViewFactory() *EventViewFactory {
	return &EventViewFactory{}
}

func (f *EventViewFactory) CreateEventViewCommand(eventId uuid.UUID, userId *string) *model.EventView {
	return &model.EventView{
		Id:	  uuid.New(),
		EventId: eventId,
		UserId:  userId,
	}
}