package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/response"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	fileUtils "github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EventHandler struct{
	eventService *service.EventService
}

func (e *EventHandler) SetupRoutes(v1Protected, v1Public *gin.RouterGroup) {
	v1Protected.POST("/event", e.CreateEvent)
	v1Public.GET("/event", e.GetEvents)

	v1Protected.GET("/event/recommendation", e.GetUserInterestedEvents)
	v1Public.GET("/event/random", e.GetRandomCategorizedEvents)

	v1Public.GET("/event/:eventId", e.GetEventDetails)
	v1Public.GET("/campus/:campusId/events", e.GetCampusEvents)

	v1Public.GET("/event/search", e.SearchEvents)
	v1Public.GET("/event/all", e.GetAllEvents)

	v1Protected.GET("/event/favorite", e.GetUserFavoritedEvents)
}

func NewEventHandler(db *gorm.DB, redis *redis.Client, client *asynq.Client) types.Route {
	return &EventHandler{
		eventService: service.NewEventService(db, redis, client),
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

	var userId *string = nil
	userQuery := c.Query("user")
	if userQuery != "" {
		userId = &userQuery
	}

	zap.L().Sugar().Infof("Fetching events with limit: %d, userId: %v", limit, userId)

	events, err := e.eventService.GetEvents(&command.GetEventsCommand{Limit: limit, UserId: userId})
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

func (e *EventHandler) GetUserInterestedEvents(c *gin.Context) {
	userId, _ := c.Get("sub")

	result, err := e.eventService.GetUserInterestedEvents(
		c.Request.Context(),
		&command.GetUserInterestedEventsCommand{
			UserId: userId.(string),
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: result.Result,
	})
}

func (e *EventHandler) GetRandomCategorizedEvents(c *gin.Context) {
	result, err := e.eventService.GetRandomEvents(c.Request.Context(), &command.GetRandomEventsCommand{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data:    result.Result,
	})
}

func (e *EventHandler) GetEventDetails(c *gin.Context) {
	eventIdStr := c.Param("eventId")
	if eventIdStr == "" {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "eventId is required",
		})
		return
	}

	eventId, err := uuid.Parse(eventIdStr)
	if err != nil || eventId == uuid.Nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidUUID.Error(),
		})
		return
	}

	var userId *string = nil

	if userQuery := c.Query("user"); userQuery != "" {
		userId = &userQuery
	}

	com := &command.GetEventDetailsCommand{
		EventID: eventId,
		UserId:  userId,
	}

	result, err := e.eventService.GetEventDetails(com)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data:    result.Result,
	})
}

func (e *EventHandler) GetCampusEvents(c *gin.Context) {
	campusIdStr := c.Param("campusId")
	if campusIdStr == "" {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "campusId is required",
		})
		return
	}

	campusId, err := uuid.Parse(campusIdStr)
	if err != nil || campusId == uuid.Nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrInvalidUUID.Error(),
		})
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "invalid page number",
		})
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if limit <= 0 {
		limit = 8
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "invalid limit number",
		})
		return
	}

	var userId *string = nil

	if userQuery := c.Query("user"); userQuery != "" {
		userId = &userQuery
	}

	com := &command.GetCampusEventsCommand{
		CampusId:   campusId,
		Limit: 	limit,
		Page: 	page,
		UserId: userId,
	}

	result, err := e.eventService.GetEventsByCampus(com)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: response.PaginationResponse{
			Rows:       result.Result,
			TotalRows:  result.TotalRows,
			TotalPages: result.TotalPages,
			Page:       result.Page,
			Limit:      result.Limit,
		}  ,
	})
}


func (e *EventHandler) SearchEvents(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "query parameter is required",
		})
		return
	}

	events, err := e.eventService.SearchEvents(&command.GetSearchEventCommand{Query: query})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data:    events.Result,
	})
}

func (e *EventHandler) GetAllEvents(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "invalid page number",
		})
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if limit <= 0 {
		limit = 8
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "invalid limit number",
		})
		return
	}

	command := &command.GetAllEventCommand{
		Limit: limit,
		Page: page,
	}

	query := c.Query("query")
	if query != "" {
		command.Filters = append(command.Filters, "title ILIKE ?")
		command.FilterArgs = append(command.FilterArgs, []interface{}{"%" + query + "%"})
	}

	price := c.Query("price")
	if price != "" {
		command.Filters = append(command.Filters, "ticket_type = ?")
		if price == "free" {
			command.FilterArgs = append(command.FilterArgs, []interface{}{"free"})
		} else if price == "paid" {
			command.FilterArgs = append(command.FilterArgs, []interface{}{"ticketed"})
		} else {
			c.JSON(http.StatusBadRequest, &types.ErrorResponse{
				Error: "invalid price filter",
			})
			return
		}
	}

	date := c.Query("date")
	if date != "" {
		if date == "today" {
			startOfDay := time.Now().Truncate(24 * time.Hour)
			endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Nanosecond)
			command.Filters = append(command.Filters, "start_time BETWEEN ? AND ?")
			command.FilterArgs = append(command.FilterArgs, []interface{}{startOfDay, endOfDay})
		} else if date == "tomorrow" {
			startOfTomorrow := time.Now().Truncate(24 * time.Hour).AddDate(0, 0, 1)
			endOfTomorrow := startOfTomorrow.Add(24 * time.Hour).Add(-time.Nanosecond)
			command.Filters = append(command.Filters, "start_time BETWEEN ? AND ?")
			command.FilterArgs = append(command.FilterArgs, []interface{}{startOfTomorrow, endOfTomorrow})
		} else if date == "this-week" {
			now := time.Now()
			weekday := int(now.Weekday())
			if weekday == 0 {
				weekday = 7
			}
			startOfWeek := now.Truncate(24 * time.Hour).AddDate(0, 0, -weekday+1)
			endOfWeek := startOfWeek.AddDate(0, 0, 7).Add(-time.Nanosecond)
			command.Filters = append(command.Filters, "start_time BETWEEN ? AND ?")
			command.FilterArgs = append(command.FilterArgs, []interface{}{startOfWeek, endOfWeek})
		} else if date == "this-month" {
			now := time.Now()
			startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
			command.Filters = append(command.Filters, "start_time BETWEEN ? AND ?")
			command.FilterArgs = append(command.FilterArgs, []interface{}{startOfMonth, endOfMonth})
		} else if date == "this-year" {
			now := time.Now()
			startOfYear := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
			endOfYear := startOfYear.AddDate(1, 0, 0).Add(-time.Nanosecond)
			command.Filters = append(command.Filters, "start_time BETWEEN ? AND ?")
			command.FilterArgs = append(command.FilterArgs, []interface{}{startOfYear, endOfYear})
		} else {
			c.JSON(http.StatusBadRequest, &types.ErrorResponse{
				Error: "invalid date filter",
			})
			return
		}
	}

	categories := c.Query("categories")
	if categories != "" {
		command.Filters = append(command.Filters, "category_id IN (?)")
		idsInterface := make([]interface{}, 0)
		splittedCategories := strings.Split(categories, ",")
		for _, category := range splittedCategories {
			categoryId, err := uuid.Parse(strings.TrimSpace(category))
			if err != nil || categoryId == uuid.Nil {
				c.JSON(http.StatusBadRequest, &types.ErrorResponse{
					Error: errors.ErrInvalidUUID.Error(),
				})
				return
			}
			idsInterface = append(idsInterface, categoryId)
		}

		zap.L().Sugar().Infof("Categories filter (raw): %v", categories)
		zap.L().Sugar().Infof("Categories filter (parsed): %v", idsInterface)

		command.FilterArgs = append(command.FilterArgs, []interface{}{idsInterface})
	}

	result, err := e.eventService.GetAllEvents(command)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: &response.PaginationResponse{
			Rows: result.Result,
			TotalRows: int64(result.TotalRows),
			TotalPages: result.TotalPage,
			Page: result.Page,
			Limit: result.Limit,
		},
	}) 
}

func (e *EventHandler) GetUserFavoritedEvents(c *gin.Context) {
	userId, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: errors.ErrMissingFields.Error(),
		})
		return
	}

	results, err := e.eventService.GetUserFavoritedEvent(&command.GetUserFavoritedEventCommand{
		UserId: userId.(string),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: results.Result,
	})
}