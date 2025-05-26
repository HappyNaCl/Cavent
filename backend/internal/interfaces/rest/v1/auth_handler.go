package v1

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/request"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/dto/response"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth/gothic"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *service.AuthService
	userService *service.UserService
}

func NewAuthRoute(db *gorm.DB, redis *redis.Client) types.Route {
	return &AuthHandler{
		authService: service.NewAuthService(db, redis),
		userService: service.NewUserService(db, redis),
	}
}

func (a *AuthHandler) SetupRoutes(v1 *gin.RouterGroup) {
	v1.POST("/auth/register", a.registerUser)
	v1.POST("/auth/login", a.loginUser)

	v1.GET("/auth/:provider", a.loginOAuthUser)
	v1.GET("/auth/:provider/callback", a.handleOAuthCallback)

	v1.GET("/auth/refresh", a.refresh)

	v1.Use(AuthMiddleware())
	v1.GET("/auth/me", a.checkMe)
	v1.GET("/auth/logout", a.logout)
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

	req.Email = strings.TrimSpace(req.Email)
	req.Email = strings.ToLower(req.Email)
	req.Name = strings.TrimSpace(req.Name)

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

	c.JSON(http.StatusCreated, &types.SuccessResponse{
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
// @Success 201 {string} string "Redirects to OAuth provider for authentication"
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
// @Success 302 {string} string "Redirects to frontend with access token as query parameter"
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

	frontendUrl := os.Getenv("FRONTEND_URL")

	c.Redirect(http.StatusFound, frontendUrl + fmt.Sprintf("?token=%s", accessToken))
}

// CheckMe godoc
// @Summary Check current user
// @Description Get the current user's information
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} types.SuccessResponse{data=response.CheckMeResponse}
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /auth/me [get]
func (a *AuthHandler) checkMe(c *gin.Context){
	userId, ok := c.Get("sub")
	if !ok {
		c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
			Error: "unauthorized",
		})
		c.Abort()
		return
	}

	command := &command.GetBriefUserCommand{
		UserId: userId.(string),
	}

	res, err := a.userService.GetBriefUser(command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user := res.Result
	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: &response.CheckMeResponse{
			User: user,
		},
	})
}

// Refresh godoc
// @Summary Refresh access token
// @Description Refresh the access token using the refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} types.SuccessResponse{data=response.RefreshResponse}
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /auth/refresh [get]
func (a *AuthHandler) refresh(c *gin.Context) {
	appDomain := os.Getenv("APP_DOMAIN")
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
			Error: "unauthorized",
		})
		return
	}

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("REFRESH_JWT_SECRET")), nil
		},
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
			Error: "invalid or expired refresh token",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
				Error: "invalid user ID in refresh token",
			})
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok || exp < float64(time.Now().Unix()) {
			zap.L().Sugar().Debugf("%d %d", exp, time.Now().Unix())
			c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
				Error: "refresh token expired",
			})
			return
		}

		refreshTokenFactory := factory.NewRefreshTokenFactory()
		newRefreshToken, err := refreshTokenFactory.GetRefreshToken(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		res, err := a.userService.GetBriefUser(&command.GetBriefUserCommand{
			UserId: userId,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		userResult := res.Result

		accessTokenFactory := factory.NewAccessTokenFactory()
		newAccessToken, err := accessTokenFactory.GetAccessToken(userResult.Id, userResult.Email, userResult.Name, userResult.FirstTimeLogin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, &types.ErrorResponse{
				Error: err.Error(),
			})
			return
		}


		expireTime := 3600 * 24 * 30 // 30 days

		c.SetCookie("refresh_token", newRefreshToken, expireTime, "/", appDomain, false, true)
		c.JSON(http.StatusOK, &types.SuccessResponse{
			Message: "success",
			Data: &response.RefreshResponse{
				AccessToken:  newAccessToken,
			},
		})
	} else {
		c.SetCookie("refresh_token", "", -1, "/", appDomain, false, true)
		c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
			Error: "invalid refresh token",
		})
	}
}

// Logout godoc
// @Summary Logout user
// @Description Logout user by clearing the refresh token cookie
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} types.SuccessResponse{data=interface{}}
// @Failure 500 {object} types.ErrorResponse
// @Router /auth/logout [get]	
func (a *AuthHandler) logout(c *gin.Context) {
	appDomain := os.Getenv("APP_DOMAIN")

	c.SetCookie("refresh_token", "", -1, "/", appDomain, false, true)
	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: nil,
	})
}