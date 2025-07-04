package common

import "github.com/google/uuid"

type BriefEventResult struct {
	Id       		uuid.UUID `json:"id"`
	Title    		string    `json:"title"`
	StartDate 		int64     `json:"startDate"`
	EndDate   		*int64    `json:"endDate,omitempty"`
	CampusName 		string    `json:"campusName"`
	Location 		string    `json:"location"`
	BannerUrl 		string    `json:"bannerUrl"`
	TicketType 		string    `json:"ticketType"`
	TicketPrice 	float64   `json:"ticketPrice"`
	CategoryName 	string    `json:"categoryName"`
	FavoriteCount 	int64     `json:"favoriteCount"`
	IsFavorited 	bool       `json:"isFavorited,omitempty"`
}