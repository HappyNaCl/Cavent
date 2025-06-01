package response

type PaginationResponse struct {
	Rows 	 interface{} 	`json:"rows"`
	TotalRows int64       	`json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}