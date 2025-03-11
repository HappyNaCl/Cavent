package interfaces

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Run(port int) error{
	r := setupRoutes()
	log.Printf("Server running on port %d", port)
	return r.Run(fmt.Sprintf(":%d", port));
}

func setupRoutes() *gin.Engine{
	r := gin.Default()

	r.GET("/", index)

	return r
}

func index(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

