package mapper

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

func NewEventResultFromEvent(event *model.Event) *common.EventResult {
	if event == nil {
		return nil
	}

	return &common.EventResult{
		Id:          event.Id,
		Title:       event.Title,
		CampusId:    event.CampusId,
		CreatedById: event.CreatedById,
		EventType:   event.EventType,
		TicketType:  event.TicketType,
		StartTime:   event.StartTime,
		Location:    event.Location,
		BannerUrl:   event.BannerUrl,
		EndTime:   	 event.EndTime,
		Description: event.Description,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
}