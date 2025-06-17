package v1

import (
	"errors"
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	domainerrors "github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	fileUtils "github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/utils"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(db *gorm.DB, redis *redis.Client, asynq *asynq.Client) types.Route {
	userService := service.NewUserService(db, redis, asynq)
	return &UserHandler{
		userService: userService,
	}
}

func (u UserHandler) SetupRoutes(v1Protected, v1Public *gin.RouterGroup) {
	v1Protected.PUT("/user/interest", u.UpdateUserInterest)
	v1Protected.GET("/user/interest", u.GetUserInterests)

	v1Protected.PUT("/user/campus", u.UpdateUserCampus)

	v1Protected.GET("/user/profile", u.GetUserProfile)
	v1Protected.PUT("/user/profile", u.UpdateUserProfile)
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
			Error: domainerrors.ErrMissingFields.Error(),
		})
		return
	}

	userId, ok := c.Get("sub")
	if !ok {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: domainerrors.ErrMissingFields.Error(),
		})
	}

	req.UserId = userId.(string)

	if len(req.CategoryIds) <= 0 {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: domainerrors.ErrInterestLength.Error(),
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
			Error: domainerrors.ErrMissingFields.Error(),
		})
		return
	}

	result, err := u.userService.GetUserInterests(c.Request.Context(), &command.GetUserInterestsCommand{
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
		Data:    result.Result,
	})
}

func (u UserHandler) UpdateUserCampus(c *gin.Context) {
	var req request.UpdateUserCampusRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: domainerrors.ErrMissingFields.Error(),
		})
		return
	}

	userId, ok := c.Get("sub")
	if !ok {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: domainerrors.ErrMissingFields.Error(),
		})
		return
	}

	req.UserId = userId.(string)

	zap.L().Sugar().Infof("userId: %s, inviteCode: %s", req.UserId, req.InviteCode)

	res, err := u.userService.UpdateUserCampus(req.ToCommand())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{
		Message: "success",
		Data:    res.User,
	})
}

func (u UserHandler) GetUserProfile(c *gin.Context) {
	userId := c.GetString("sub")

	result, err := u.userService.GetUserProfile(c.Request.Context(), &command.GetUserProfileCommand{
		UserId: userId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{
		Message: "success",
		Data:    result.Result,
	})
}

func (u UserHandler) UpdateUserProfile(c *gin.Context) {
	var req request.UpdateUserProfileRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: domainerrors.ErrMissingFields.Error(),
		})
		return
	}

	userId := c.GetString("sub")
	if userId == "" {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error: domainerrors.ErrMissingFields.Error(),
		})
		return
	}

	file, header, err := c.Request.FormFile("profileImage")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			c.JSON(http.StatusBadRequest, &types.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
	}

	command := req.ToCommand()
	var fileBytes []byte
	var fileExt string

	command.UserId = userId

	if file != nil && header != nil {
		fileBytes, fileExt, err = fileUtils.ReadMultipartFile(file, header)
		if err != nil {
			c.JSON(http.StatusBadRequest, &types.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		command.AvatarBytes = fileBytes
		command.AvatarExt = &fileExt
	}

	res, err := u.userService.UpdateUserProfile(c.Request.Context(), command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{
		Message: "success",
		Data:    res.Result,
	})
}