package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, redis *redis.Client, client *asynq.Client) *gin.Engine  {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	v1Protected := r.Group("/api/v1")
	v1Protected.Use(AuthMiddleware())

	v1Public := r.Group("/api/v1")

	authRoute := NewAuthRoute(db, redis, client)
	authRoute.SetupRoutes(v1Protected, v1Public)

	categoryRoute := NewCategoryHandler(db, redis)
	categoryRoute.SetupRoutes(v1Protected, v1Public)

	userRoute := NewUserHandler(db, redis, client)
	userRoute.SetupRoutes(v1Protected, v1Public)
	
	eventRoute := NewEventHandler(db, redis, client)
	eventRoute.SetupRoutes(v1Protected, v1Public)
	
	favRoute := NewFavoriteHandler(db, redis)
	favRoute.SetupRoutes(v1Protected, v1Public)

	campusHandler := NewCampusHandler(db)
	campusHandler.SetupRoutes(v1Protected, v1Public)

	ticketTypeHandler := NewTicketTypeHandler(db)
	ticketTypeHandler.SetupRoutes(v1Protected, v1Public)

	transactionHandler := NewTransactionHandler(db, client)
	transactionHandler.SetupRoutes(v1Protected, v1Public)
	
	ticketHandler := NewTicketHandler(db)
	ticketHandler.SetupRoutes(v1Protected, v1Public)
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}