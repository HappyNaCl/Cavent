package v1

import (
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

// UpdateUserInterest godoc
// @Summary      Update user interests
// @Description  Update user interests
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Bearer token for authentication"
// @Param        request  body  request.UpdateUserInterestRequest  true  "Update user interests request"
// @Success      200  {object}  types.SuccessResponse
// @Failure      400  {object}  types.ErrorResponse
// @Failure      500  {object}  types.ErrorResponse
// @Router       /user/interest [put]
// @Security     BearerAuth
func (u UserHandler) UpdateUserInterest(c *gin.Context) {
	var req request.UpdateUserInterestRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: errors.ErrMissingFields.Error(),
		})
		return
	}

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

	var token string

	firstTimeLogin, exists := c.Get("firstTimeLogin")
	if exists && firstTimeLogin.(bool) {
		id, _ := c.Get("sub")
        email, _ := c.Get("email")
		role, _ := c.Get("role")


		accessTokenFactory := factory.NewAccessTokenFactory()
		token, err = accessTokenFactory.GetAccessToken(id.(string), email.(string), role.(string), false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		userId, _ := c.Get("sub")
		user, err := u.userService.GetBriefUser(&command.GetBriefUserCommand{
			UserId: userId.(string),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, types.SuccessResponse{
			Message: "success",
			Data: gin.H{
				"accessToken": token,
				"user":        user.Result,
			},
		})
		return
	}
	

	c.JSON(http.StatusOK, types.SuccessResponse{
		Message: "success",
		Data:    nil,
	})
}

// GetUserInterests godoc
// @Summary      Get user interests
// @Description  Get user interests
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Bearer token for authentication"
// @Success      200  {object}  types.SuccessResponse
// @Failure      400  {object}  types.ErrorResponse
// @Failure      500  {object}  types.ErrorResponse
// @Router       /user/interest [get]
// @Security     BearerAuth
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