package v1

import (
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
	"go.uber.org/zap"
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
	var req *request.CreateEventRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrMissingFields.Error(),
		})
		return
	}
	
	if req.EventType != "single" && req.EventType != "recurring" {
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

	if time.Unix(req.StartTime, 0).Before(now) {
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

	for _, ticket := range req.Ticket {
		com.Ticket = append(com.Ticket, common.TicketResult{
			Name:     ticket.Name,
			Price:    ticket.Price,
			Quantity: ticket.Quantity,
		})
	}

	userId, _ := c.Get("sub")
	zap.L().Sugar().Infof("User ID: %v", userId)
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
