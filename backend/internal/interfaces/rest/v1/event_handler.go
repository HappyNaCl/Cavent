package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/utils"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type EventHandler struct{
	eventService *service.EventService
}

func (e *EventHandler) SetupRoutes(v1 *gin.RouterGroup) {
	v1.POST("/event", e.CreateEvent)
}

func NewEventHandler(db *gorm.DB, client *asynq.Client) types.Route {
	return &EventHandler{
		eventService: service.NewEventService(db, client),
	}
}

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
	if req.Ticket != nil  {
		if err := json.Unmarshal([]byte(*req.Ticket), &tickets); err != nil {
			c.JSON(http.StatusBadRequest, &types.ErrorResponse{
				Error: "invalid ticket format: " + err.Error(),
			})
			return
		}
	}

	fileBytes, fileExt, err := utils.ReadMultipartFile(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	com := &command.CreateEventCommand{
		Title:       req.Title,
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
