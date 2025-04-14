package interfaces

import (
	"fmt"
	"log"

	"github.com/HappyNaCl/Cavent/backend/interfaces/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(port int) error{
	r := setupRoutes()
	log.Printf("Server running on port %d", port)
	return r.Run(fmt.Sprintf(":%d", port));
}

func setupRoutes() *gin.Engine{
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	authHandler := handler.AuthHandler{}
	preferenceHandler := handler.NewPreferenceHandler()
	tagHandler := handler.NewTagHandler()
	userHandler := handler.NewUserHandler()

	r.GET("/", index)

	auth := r.Group("/api/auth")
	auth.Use(UnauthMiddleware())
	auth.POST("/register", authHandler.RegisterUser)
	auth.POST("/login", authHandler.LoginCredential)
	auth.GET("/:provider", authHandler.LoginWithOAuth)
	auth.GET("/:provider/callback", authHandler.LoginWithOAuthCallback)

	authProtected := r.Group("/api/auth")
	authProtected.Use(AuthMiddleware())
	authProtected.POST("/logout", authHandler.Logout)
	authProtected.GET("/me", authHandler.CheckMe)

	tags := r.Group("/api/tags")
	tags.Use(AuthMiddleware())
	tags.GET("", tagHandler.GetAllTagsWithType)

	profile := r.Group("/api/user")
	profile.Use(AuthMiddleware())
	profile.GET("/tag", userHandler.GetUserTag)
	profile.PUT("/preference", preferenceHandler.UpdatePreferences)

	protected := r.Group("/api/protected")
	protected.Use(AuthMiddleware())
	protected.GET("/", protectedIndex)
	return r
}

func index(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func protectedIndex(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Protected",
	})
}

