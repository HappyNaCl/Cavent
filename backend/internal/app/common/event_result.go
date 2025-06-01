package common

import (
	"time"

	"github.com/google/uuid"
)

type EventResult struct {
	Id     	 	uuid.UUID 	`json:"id"`
	Title  	 	string 		`json:"title"`
	CreatedById string 		`json:"createdById"`
	CampusId 	uuid.UUID 	`json:"campusId"`
	EventType 	string 		`json:"eventType"`
	TicketType 	string 		`json:"ticketType"`
	StartTime 	int64  	`json:"startTime"`
	EndTime  	*int64  `json:"endTime"`
	Location 	string 	   	`json:"location"`
	Description *string 	`json:"description,omitempty"`
	BannerUrl 	string 		`json:"bannerUrl"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	IsFavorited bool       `json:"isFavorited,omitempty"`
	Tickets     []*TicketResult `json:"tickets,omitempty"`
	Campus      BriefCampusResult `json:"campus"`
}