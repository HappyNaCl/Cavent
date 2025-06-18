package v1

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TicketHandler struct {
	service *service.TicketService
}

// SetupRoutes implements types.Route.
func (t *TicketHandler) SetupRoutes(v1Protected *gin.RouterGroup, v1Public *gin.RouterGroup) {
	v1Protected.GET("/user/tickets", t.GetUserTickets)
}

func NewTicketHandler(db *gorm.DB) types.Route {
	return &TicketHandler{
		service: service.NewTicketService(db),
	}
}

func (t *TicketHandler) GetUserTickets(c *gin.Context) {
	userId := c.GetString("sub")

	cmd := &command.GetUserTicketsCommand{
		UserId: userId,
	}

	result, err := t.service.GetUserTickets(c, cmd)
	if err != nil {
		c.JSON(500, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(200, &types.SuccessResponse{
		Message: "success",
		Data:    result.Tickets,
	})
}
