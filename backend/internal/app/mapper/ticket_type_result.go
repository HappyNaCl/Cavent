package mapper

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

func NewTicketTypeResultFromTicketType(ticketType []*model.TicketType) []*common.TicketTypeResult {
	var ticketTypeResults []*common.TicketTypeResult
	for _, t := range ticketType {
		ticketTypeResults = append(ticketTypeResults, &common.TicketTypeResult{
			Id:          t.Id,
			Name:        t.Name,
			Price:       t.Price,
			Quantity:    t.Quantity,
			Sold: 		 t.Sold,
		})
	}
	return ticketTypeResults
}