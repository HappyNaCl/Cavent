package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	fileUtils "github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type EventHandler struct{
	eventService *service.EventService
}

func (e *EventHandler) SetupRoutes(v1 *gin.RouterGroup) {
	v1.POST("/event", e.CreateEvent)
	v1.GET("/event", e.GetEvents)
}

func NewEventHandler(db *gorm.DB, client *asynq.Client) types.Route {
	return &EventHandler{
		eventService: service.NewEventService(db, client),
	}
}

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event with details and optional tickets
// @Tags Event
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Event title"
// @Param event_type formData string true "Event type (single or recurring)"
// @Param ticket_type formData string true "Ticket type (ticketed or free)"
// @Param start_time formData int64 true "Event start time in Unix timestamp"
// @Param end_time formData int64 false "Event end time in Unix timestamp (optional, must be same day as start time)"
// @Param location formData string true "Event location"
// @Param description formData string false "Event description (optional)"
// @Param banner formData file true "Event banner image"
// @Param ticket formData string false "JSON string of tickets (optional, required if ticket_type is ticketed)"
// @Success 201 {object} types.SuccessResponse{data=common.EventResult} "Event created successfully"
// @Failure 400 {object} types.ErrorResponse "Bad request"
// @Failure 500 {object} types.ErrorResponse "Internal server error"
// @Router /v1/event [post]
func (e *EventHandler) CreateEvent(c *gin.Context) {
	var req request.CreateEventRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	
	if req.EventType != "single" && req.EventType != "recurring" {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidEventType.Error(),
		})
		return
	}

	if req.TicketType != "ticketed" && req.TicketType != "free" {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidTicketType.Error(),
		})
		return
	}

	now := time.Now()
	startTime := time.Unix(req.StartTime, 0)

	if startTime.Before(now) {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidStartTime.Error(),
		})
		return
	}

	if req.EndTime != nil {
		endTime := time.Unix(*req.EndTime, 0)
		if endTime.Before(startTime) {
			c.JSON(http.StatusBadRequest, &types.ErrorResponse{
				Error: errors.ErrInvalidEndTime.Error(),
			})
			return
		}

		if endTime.Day() != startTime.Day() {
			c.JSON(http.StatusBadRequest, &types.ErrorResponse{
				Error: errors.ErrEndDateMustBeSameDay.Error(),
			})
			return
		}
		
	}

	file, header, err := c.Request.FormFile("banner")
	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrMissingBanner.Error(),
		})
		return
	}

	var tickets []common.TicketResult
	if req.Ticket != nil && req.TicketType == "ticketed" {
		if err := json.Unmarshal([]byte(*req.Ticket), &tickets); err != nil {
			c.JSON(http.StatusBadRequest, &types.ErrorResponse{
				Error: "invalid ticket format: " + err.Error(),
			})
			return
		}
	}

	fileBytes, fileExt, err := fileUtils.ReadMultipartFile(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	categoryUUID, err := uuid.Parse(req.CategoryId)
	if err != nil || categoryUUID == uuid.Nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidUUID.Error(),
		})
		return
	}

	com := &command.CreateEventCommand{
		Title:       req.Title,
		CategoryId:  categoryUUID,
		EventType:   req.EventType,
		TicketType:  req.TicketType,
		StartTime:   time.Unix(req.StartTime, 0),
		Location:   req.Location,
		BannerBytes: fileBytes,
		BannerExt:   fileExt,
	}

	if req.EndTime != nil {
		endTime := time.Unix(*req.EndTime, 0)
		com.EndTime = &endTime
	} else {
		com.EndTime = nil
	}

	if req.Description != nil {
		com.Description = req.Description
	} else {
		com.Description = nil
	}

	for _, ticket := range tickets {
		com.Ticket = append(com.Ticket, common.TicketResult{
			Name:     ticket.Name,
			Price:    ticket.Price,
			Quantity: ticket.Quantity,
		})
	}

	userId, _ := c.Get("sub")
	com.CreatedById = userId.(string)	

	result, err := e.eventService.CreateEvent(com)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &types.SuccessResponse{
		Message: "success",
		Data: result.Result,
	})
}

// GetEvents godoc
// @Summary Get all events
// @Description Get a list of events with pagination
// @Tags Event
// @Accept json
// @Produce json
// @Param limit query int false "Number of events to return (default is 20)"
// @Success 200 {object} types.SuccessResponse{data=[]common.BriefEventResult} "List of events"
// @Failure 500 {object} types.ErrorResponse "Internal server error"
// @Router /v1/event [get]
func (e *EventHandler) GetEvents(c *gin.Context) {
	limit := 20

	if limitStr := c.Query("limit"); limitStr != "" {
		if limitValue, err := strconv.Atoi(limitStr); err == nil && limitValue > 0 {
			limit = limitValue
		} 
	}


	events, err := e.eventService.GetEvents(&command.GetEventsCommand{Limit: limit})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: events.Result,
	})
}