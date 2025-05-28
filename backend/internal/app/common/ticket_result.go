package common

type TicketResult struct {
	Name 	  string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}