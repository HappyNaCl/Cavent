package v1

import (
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CampusHandler struct {
	campusService *service.CampusService
}

func (c *CampusHandler) SetupRoutes(v1Protected *gin.RouterGroup, v1Public *gin.RouterGroup) {
	v1Public.GET("/campus", c.getAllCampus)
}

func (ch *CampusHandler) getAllCampus(c *gin.Context) {
	result, err := ch.campusService.GetAllCampus(&command.GetAllCampusCommand{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(200, &types.SuccessResponse{
		Message: "success",
		Data:    result.Result,
	})
}

func NewCampusHandler(db *gorm.DB) types.Route {
	return &CampusHandler{
		campusService: service.NewCampusService(db),
	}
}
