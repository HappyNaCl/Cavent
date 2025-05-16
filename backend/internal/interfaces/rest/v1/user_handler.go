package v1

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app"
	"github.com/HappyNaCl/Cavent/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService *app.UserService
	logger *zap.Logger
}

func NewUserHandler(db *gorm.DB, redis *redis.Client, logger *zap.Logger) types.Route {
	userService := app.NewUserService(db, redis, logger)
	return &UserHandler{
		userService: userService,
		logger: logger,
	}
}

func (u UserHandler) SetupRoutes(r *gin.RouterGroup) {

}