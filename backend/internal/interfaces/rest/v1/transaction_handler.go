package v1

import (
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	service *service.TransactionService
	asynq *asynq.Client
}

// SetupRoutes implements types.Route.
func (t *TransactionHandler) SetupRoutes(v1Protected *gin.RouterGroup, v1Public *gin.RouterGroup) {
	v1Protected.POST("/checkout", t.Checkout)
}

func NewTransactionHandler(db *gorm.DB, asynq *asynq.Client) types.Route {
	return &TransactionHandler{
		service: service.NewTransactionService(db, asynq),
		asynq: asynq,
	}
}

func (t *TransactionHandler) Checkout(c *gin.Context) {
	var req *request.CheckoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Sugar().Debugf("Checkout request binding error: %v", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	userId := c.GetString("sub")
	eventUuid, err := uuid.Parse(req.EventId)
	if err != nil {
		zap.L().Sugar().Debugf("Invalid event UUID: %v", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	com, err := req.ToCommand()
	if err != nil {
		zap.L().Sugar().Debugf("Error converting request to command: %v", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	
	com.UserId = userId
	com.EventId = eventUuid

	_, err = t.service.Checkout(c.Request.Context(), com)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}


}
