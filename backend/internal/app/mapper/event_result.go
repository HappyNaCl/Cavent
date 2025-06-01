package mapper

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

func NewEventResultFromEvent(event *model.Event, isFavorited bool) *common.EventResult {
	if event == nil {
		return nil
	}
	
	result := &common.EventResult{
		Id:          event.Id,
		Title:       event.Title,
		CampusId:    event.CampusId,
		CreatedById: event.CreatedById,
		EventType:   event.EventType,
		TicketType:  event.TicketType,
		StartTime:   event.StartTime.Unix(),
		Location:    event.Location,
		BannerUrl:   event.BannerUrl,
		Description: event.Description,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
		IsFavorited: isFavorited,
	}

	if event.TicketType == "ticketed" {
		tickets := make([]*common.TicketResult, 0, len(event.TicketTypes))
		for _, ticket := range event.TicketTypes {
			tickets = append(tickets, &common.TicketResult{
				Id:          ticket.Id,
				Name:        ticket.Name,
				Price:       ticket.Price,
				Quantity:    ticket.Quantity,
			})
		}
		result.Tickets = tickets
	}

	if event.EndTime != nil {
		endTime := event.EndTime.Unix()
		result.EndTime = &endTime
	} else {
		result.EndTime = nil
	}
	
	result.Campus = common.BriefCampusResult{
		Id:   event.Campus.Id,
		Name: event.Campus.Name,
		ProfileUrl: event.Campus.LogoUrl,
	}

	return result
}

func NewBriefEventResultFromEvent(event *model.Event, isFavorited bool) *common.BriefEventResult {
	if event == nil {
		return nil
	}

	cheapPrice := getCheapestTicketPrice(event.TicketTypes)
	var endDate *int64
	if event.EndTime != nil {
		endDate = new(int64)
		*endDate = event.EndTime.Unix()
	} else {
		endDate = nil
	}

	return &common.BriefEventResult{
		Id:          event.Id,
		Title:       event.Title,
		StartDate: 	 event.StartTime.Unix(),
		EndDate:  	 endDate,
        CampusName:  event.Campus.Name,
		TicketType:  event.TicketType,
		TicketPrice: cheapPrice,
		Location:    event.Location,
		BannerUrl:   event.BannerUrl,
		CategoryName: event.Category.Name,
		FavoriteCount: event.FavoriteCount,
		IsFavorited: isFavorited,
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