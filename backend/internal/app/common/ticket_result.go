package common

import "github.com/google/uuid"

type TicketResult struct {
	Id 		  uuid.UUID  `json:"id"`  
	Name 	  string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}