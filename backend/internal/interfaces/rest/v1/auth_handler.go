package v1

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app"
	"github.com/HappyNaCl/Cavent/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *app.AuthService
	logger *zap.Logger
}

func NewAuthRoute(db *gorm.DB, redis *redis.Client, logger *zap.Logger) types.Route {
	return &AuthHandler{
		authService: app.NewAuthService(db, redis, logger),
		logger: logger,
	}
}

func (a *AuthHandler) SetupRoutes(r *gin.RouterGroup) {

}