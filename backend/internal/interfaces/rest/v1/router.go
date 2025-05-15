package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine  {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	return r
}