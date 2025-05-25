package v1

import (
	"net/http"
	"os"

	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/response"
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
	v1.POST("/auth/login", a.loginUser)
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
	refreshToken, err := refreshTokenFactory.GetRefreshToken(userResult.Id, userResult.Email)
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
		c.Abort()
		return
	}

	res, err := a.authService.LoginUser(req.ToCommand())
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

	expireTime := 3600 * 24
	if req.RememberMe {
		expireTime = 3600 * 24 * 30
	}

	c.SetCookie("refresh_token", refreshToken, expireTime, "/", appDomain, false, true)

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: &response.LoginResponse{
			AccessToken:  accessToken,
		},
	})
}

func (a *AuthHandler) loginOAuthUser(c *gin.Context) {
	// gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// }

// 	user, err := application.RegisterOrLoginOauthUser(gothUser, gothUser.Provider)
// 	if err != nil || user == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}

// 	token, err := application.GenerateJWT(application.JWTClaims{
// 		Id: user.Id,
// 		Provider: user.Provider,
// 		Email: user.Email,
// 		AvatarUrl: user.AvatarUrl,
// 		Name: user.Name,
// 		FirstTimeLogin: user.FirstTimeLogin,
// 		Exp: 3600*24,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}

// 	appDomain := os.Getenv("APP_DOMAIN")
// 	c.SetCookie("token", token, 3600*24, "/", appDomain, false, true)
}