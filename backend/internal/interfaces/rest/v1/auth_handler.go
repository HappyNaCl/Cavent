package v1

import (
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthRoute(db *gorm.DB, redis *redis.Client) types.Route {
	return &AuthHandler{
		authService: service.NewAuthService(db, redis),
	}
}

func (a *AuthHandler) SetupRoutes(v1 *gin.RouterGroup) {
	v1.POST("/auth/register", a.registerUser)
}

func (a *AuthHandler) registerUser(c *gin.Context) {
	var req *request.RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "invalid form requests",
		})
		c.Abort()
		return
	}

	res, err := a.authService.RegisterUser(req.ToCommand())
	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: err.Error(),
		})
		c.Abort()
		return
	}

	userResult := res.Result
	accessTokenFactory := factory.NewAccessTokenFactory()
	accessToken, err := accessTokenFactory.GetAccessToken(userResult.Id, userResult.Email, userResult.Name, userResult.FirstTimeLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		c.Abort()
		return
	}

	refreshTokenFactory := factory.NewRefreshTokenFactory()
	refreshToken, err := refreshTokenFactory.GetRefreshToken(userResult.Id, userResult.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "registered successfully",
		Data: gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}
