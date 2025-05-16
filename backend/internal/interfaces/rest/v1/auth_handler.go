package v1

import (
	"github.com/HappyNaCl/Cavent/backend/internal/types"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	// authService *app.AuthService
}

func NewAuthRoute() types.Route {
	return &AuthHandler{
		// authService: app.NewAuthService(db, redis),
	}
}

func (a *AuthHandler) SetupRoutes(r *gin.RouterGroup) {
	// r.POST("/login", a.authService.Login)
	// r.POST("/register", a.authService.Register)
	// r.GET("/logout", a.authService.Logout)
	// r.GET("/refresh", a.authService.Refresh)
}