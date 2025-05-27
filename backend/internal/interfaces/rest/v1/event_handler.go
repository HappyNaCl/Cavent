package v1

import (
	"net/http"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/utils"
	"github.com/gin-gonic/gin"
)

type EventHandler struct{}

func (e *EventHandler) SetupRoutes(v1 *gin.RouterGroup) {
	panic("unimplemented")
}

func NewEventHandler() types.Route {
	return &EventHandler{}
}

func (e *EventHandler) CreateEvent(c *gin.Context) {
	var req *request.CreateEventRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrMissingFields.Error(),
		})
		return
	}
	
	if req.EventType != "single" || req.EventType != "recurring" {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidEventType.Error(),
		})
		return
	}

	if req.TicketType != "paid" && req.TicketType != "free" {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidTicketType.Error(),
		})
		return
	}

	now := time.Now()

	if time.Unix(req.StartDate, 0).Before(now.Truncate(24 * time.Hour)) {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidStartDate.Error(),
		})
		return 
	}

	if time.Unix(req.StartDate, 0).Before(now) {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidStartDate.Error(),
		})
		return
	}

	if req.EndTime != nil && !time.Unix(*req.EndTime, 0).After(time.Unix(req.StartTime, 0)) {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidEndTime.Error(),
		})
		return
	}

	file, header, err := c.Request.FormFile("banner")
	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrMissingBanner.Error(),
		})
		return
	}

	fileBytes, fileExt, err := utils.ReadMultipartFile(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
}
