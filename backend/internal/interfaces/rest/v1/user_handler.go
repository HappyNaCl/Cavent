package v1

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
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

func (u UserHandler) SetupRoutes(r *gin.RouterGroup) {

}