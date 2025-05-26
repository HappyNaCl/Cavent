package v1

import (
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(db *gorm.DB, redis *redis.Client) types.Route {
	userService := service.NewUserService(db, redis)
	return &UserHandler{
		userService: userService,
	}
}

func (u UserHandler) SetupRoutes(v1 *gin.RouterGroup) {
	v1.PUT("/user/interest", u.UpdateUserInterest)
	v1.GET("/user/interest", u.GetUserInterests)
}

func (u UserHandler) UpdateUserInterest(c *gin.Context) {
	var req request.UpdateUserInterestRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: errors.ErrMissingFields.Error(),
		})
		return
	}

	zap.L().Sugar().Debugf("CategoryIds type: %T, value: %v", req.CategoryIds, req.CategoryIds[0])

	userId, ok := c.Get("sub")
	if !ok {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: errors.ErrMissingFields.Error(),
		})
	}

	req.UserId = userId.(string)

	if len(req.CategoryIds) <= 0 {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: errors.ErrInterestLength.Error(),
		})
		return
	}

	_, err := u.userService.UpdateUserInterests(req.ToCommand())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{
		Message: "success",
		Data:    nil,
	})
}

func (u UserHandler) GetUserInterests(c *gin.Context) {
	userId, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: errors.ErrMissingFields.Error(),
		})
		return
	}

	result, err := u.userService.GetUserInterests(&command.GetUserInterestsCommand{
		UserId: userId.(string),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{
		Message: "success",
		Data:    result.CategoryTypes,
	})
}