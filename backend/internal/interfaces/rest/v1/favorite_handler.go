package v1

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type FavoriteHandler struct {
	service *service.FavoriteService
}

// SetupRoutes implements types.Route.
func (f *FavoriteHandler) SetupRoutes(v1Protected *gin.RouterGroup, v1Public *gin.RouterGroup) {
	v1Protected.POST("/event/:eventId/favorite", f.FavoriteEvent)
	v1Protected.DELETE("/event/:eventId/favorite", f.UnfavoriteEvent)
}

func (f *FavoriteHandler) FavoriteEvent(c *gin.Context) {
	eventId := c.Param("eventId")
	userId := c.GetString("sub")

	eventUuid, err := uuid.Parse(eventId)
	if err != nil {
		c.JSON(400, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	cmd := &command.FavoriteEventCommand{
		EventId: eventUuid,
		UserId:  userId,
	}
	
	count, err := f.service.FavoriteEvent(c, cmd)
	if err != nil {
		c.JSON(500, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(200, &types.SuccessResponse{
		Message: "success",
		Data:   count.Result,
	})
}

func (f *FavoriteHandler) UnfavoriteEvent(c *gin.Context) {
	eventId := c.Param("eventId")
	userId := c.GetString("sub")

	eventUuid, err := uuid.Parse(eventId)
	if err != nil {
		c.JSON(400, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	cmd := &command.UnfavoriteEventCommand{
		EventId: eventUuid,
		UserId:  userId,
	}
	
	count, err := f.service.UnfavoriteEvent(c, cmd)
	if err != nil {
		c.JSON(500, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(200, &types.SuccessResponse{
		Message: "success",
		Data:   count.Result,
	})
}

func NewFavoriteHandler(db *gorm.DB, redis *redis.Client) types.Route {
	return &FavoriteHandler{
		service: service.NewFavoriteService(db, redis),
	}
}
