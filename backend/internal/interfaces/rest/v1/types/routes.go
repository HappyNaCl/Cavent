package types

import "github.com/gin-gonic/gin"

type Route interface {
	SetupRoutes(v1Protected, v1Public *gin.RouterGroup)
}