package v1

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TicketTypeHandler struct {
	service *service.TicketTypeService
}

func (t *TicketTypeHandler) SetupRoutes(v1Protected *gin.RouterGroup, v1Public *gin.RouterGroup) {
	v1Protected.GET("/event/:eventId/tickets", t.GetTicketTypesByEventId)
}

func (t *TicketTypeHandler) GetTicketTypesByEventId(c *gin.Context) {
	eventId := c.Param("eventId")

	cmd := &command.GetTicketTypeByEventIdCommand{
		EventId: eventId,
	}

	result, err := t.service.GetTicketTypeByEventId(c, cmd)
	if err != nil {
		c.JSON(500, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(200, &types.SuccessResponse{
		Message: "success",
		Data:    result.TicketTypes,
	})
}

func NewTicketTypeHandler(db *gorm.DB) types.Route {
	return &TicketTypeHandler{
		service: service.NewTicketTypeService(db),
	}
}
