package interfaces

import (
	"fmt"
	"log"

	"github.com/HappyNaCl/Cavent/backend/interfaces/handler"
	"github.com/gin-gonic/gin"
)

func Run(port int) error{
	r := setupRoutes()
	log.Printf("Server running on port %d", port)
	return r.Run(fmt.Sprintf(":%d", port));
}

func setupRoutes() *gin.Engine{
	r := gin.Default()

	authHandler := handler.AuthHandler{}

	r.GET("/", index)

	auth := r.Group("/api/auth")
	auth.Use(UnauthMiddleware())
	auth.POST("/register", authHandler.RegisterUser)
	auth.POST("/login", authHandler.LoginCredential)
	auth.GET("/:provider", authHandler.LoginWithOAuth)
	auth.GET("/:provider/callback", authHandler.LoginWithOAuthCallback)

	// auth.POST("/logout", AuthMiddleware(), authHandler.Logout)

	authProtected := r.Group("/api/auth")
	authProtected.Use(AuthMiddleware())
	authProtected.POST("/logout", authHandler.Logout)

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

