package common

import "github.com/google/uuid"

type CheckoutTicketResult struct {
	Id 	 	 uuid.UUID  `json:"id"`
	Quantity int     `json:"quantity"`
}