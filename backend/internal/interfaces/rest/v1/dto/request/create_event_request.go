package request

type Ticket struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Price       float64 `json:"price" form:"price" binding:"required"`
	Quantity    int    `json:"quantity" form:"quantity" binding:"required"`
}

type CreateEventRequest struct {
	Title 		string 		`json:"title" form:"title" binding:"required"`
	CategoryId 	string 		`json:"categoryId" form:"categoryId" binding:"required"`
	EventType 	string 		`json:"eventType" form:"eventType" binding:"required"`
	TicketType 	string 		`json:"ticketType" form:"ticketType" binding:"required"`
	StartTime 	int64		`json:"startTime" form:"startTime" binding:"required"`
	EndTime 	*int64 		`json:"endTime" form:"endTime"`
	Location 	string 		`json:"location" form:"location" binding:"required"`
	Description *string 	`json:"description" form:"description"`
	Ticket      *string 	`json:"tickets" form:"tickets"`
}