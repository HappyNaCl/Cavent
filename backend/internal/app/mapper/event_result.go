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

func NewBriefEventResultFromEvent(event *model.Event) *common.BriefEventResult {
	if event == nil {
		return nil
	}

	cheapPrice := getCheapestTicketPrice(event.TicketTypes)

	return &common.BriefEventResult{
		Id:          event.Id,
		Title:       event.Title,
		StartDate: 	 event.StartTime.Unix(),
        CampusName:  event.Campus.Name,
		TicketType:  event.TicketType,
		TicketPrice: cheapPrice,
		Location:    event.Location,
		BannerUrl:   event.BannerUrl,
	}
}

func getCheapestTicketPrice(tickets []model.TicketType) float64 {
	if len(tickets) == 0 {
		return 0 
	}

	min := tickets[0].Price
	for _, t := range tickets[1:] {
		if t.Price < min {
			min = t.Price
		}
	}
	return min
}