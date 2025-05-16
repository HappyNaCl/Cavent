package types

import "github.com/gin-gonic/gin"

type Route interface {
	SetupRoutes(v1 *gin.RouterGroup)
}