package v1

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app"
	"github.com/HappyNaCl/Cavent/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService *app.UserService
}

func NewUserHandler(db *gorm.DB, redis *redis.Client) types.Route {
	userService := app.NewUserService(db, redis)
	return &UserHandler{
		userService: userService,
	}
}

func (u UserHandler) SetupRoutes(r *gin.RouterGroup) {

}