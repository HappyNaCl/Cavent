package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, redis *redis.Client) *gin.Engine  {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	v1 := r.Group("/api/v1")
	
	authRoute := NewAuthRoute(db, redis)
	authRoute.SetupRoutes(v1)

	categoryRoute := NewCategoryHandler(db, redis)
	categoryRoute.SetupRoutes(v1)
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}