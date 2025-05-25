package v1

import (
	"context"
	"net/http"
	"os"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/response"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
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
	v1.POST("/auth/login", a.loginUser)

	v1.GET("/auth/:provider", a.loginOAuthUser)
	v1.GET("/auth/:provider/callback", a.handleOAuthCallback)
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with name, email, and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.RegisterRequest true "Register Request"
// @Success 200 {object} types.SuccessResponse{data=response.RegisterResponse}
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /auth/register [post]
func (a *AuthHandler) registerUser(c *gin.Context) {
	appDomain := os.Getenv("APP_DOMAIN")

	var req request.RegisterRequest
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
	refreshToken, err := refreshTokenFactory.GetRefreshToken(userResult.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		c.Abort()
		return
	}

	c.SetCookie("refresh_token", refreshToken, 3600*24, "/", appDomain, false, true)

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: &response.RegisterResponse{
			AccessToken:  accessToken,
			User: userResult.ToBrief(),
		},
	})
}

// LoginUser godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login Request"
// @Success 200 {object} types.SuccessResponse{data=response.LoginResponse}
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /auth/login [post]
func (a *AuthHandler) loginUser(c *gin.Context) {
	appDomain := os.Getenv("APP_DOMAIN")

	var req request.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "invalid form requests",
		})
		return
	}

	res, err := a.authService.LoginUser(req.ToCommand())
	if err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	userResult := res.Result
	accessTokenFactory := factory.NewAccessTokenFactory()
	accessToken, err := accessTokenFactory.GetAccessToken(userResult.Id, userResult.Email, userResult.Name, userResult.FirstTimeLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	refreshTokenFactory := factory.NewRefreshTokenFactory()
	refreshToken, err := refreshTokenFactory.GetRefreshToken(userResult.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	expireTime := 3600 * 24
	if req.RememberMe {
		expireTime = 3600 * 24 * 30
	}

	c.SetCookie("refresh_token", refreshToken, expireTime, "/", appDomain, false, true)

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: &response.LoginResponse{
			AccessToken:  accessToken,
			User: userResult.ToBrief(),
		},
	})
}


// LoginOAuthUser godoc
// @Summary Login a user with OAuth
// @Description Login a user with OAuth provider
// @Tags auth
// @Accept json
// @Produce json
// @Param provider path string true "OAuth Provider" Enums(google, github, etc.)
// @Success 201 {object} types.SuccessResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /auth/{provider} [get]
func (a *AuthHandler) loginOAuthUser(c *gin.Context) {
	provider := c.Param("provider")
	
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// OAuthCallback godoc
// @Summary OAuth Callback
// @Description Handle OAuth callback after user authentication
// @Tags auth
// @Accept json
// @Produce json
// @Param provider path string true "OAuth Provider" Enums(google, github, etc.)
// @Success 200 {object} types.SuccessResponse{data=response.LoginResponse}
// @Failure 500 {object} types.ErrorResponse
// @Router /auth/{provider}/callback [get]
func (a *AuthHandler) handleOAuthCallback(c *gin.Context) {
	appDomain := os.Getenv("APP_DOMAIN")

	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	oauthCommand := &command.HandleOAuthCommand{
		Provider: c.Param("provider"),
		User:     gothUser,
	}
	
	res, err := a.authService.HandleOAuth(oauthCommand)
	userResult := res.Result
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	accessTokenFactory := factory.NewAccessTokenFactory()
	accessToken, err := accessTokenFactory.GetAccessToken(userResult.Id, userResult.Email, userResult.Name, userResult.FirstTimeLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	refreshTokenFactory := factory.NewRefreshTokenFactory()
	refreshToken, err := refreshTokenFactory.GetRefreshToken(userResult.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	expireTime := 3600 * 24
	c.SetCookie("refresh_token", refreshToken, expireTime, "/", appDomain, false, true)

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: &response.LoginResponse{
			AccessToken:  accessToken,
		},
	})
}