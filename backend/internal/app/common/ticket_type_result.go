package common

import "github.com/google/uuid"

type TicketTypeResult struct {
	Id    	 uuid.UUID  `json:"id"`
	Name  	 string		`json:"name"`
	Price 	 float64	`json:"price"`
	Quantity int		`json:"quantity"`
	Sold 	 int		`json:"sold"`
}