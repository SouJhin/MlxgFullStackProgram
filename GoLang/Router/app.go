package router

import (
	service "server/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/log", service.Log)
	return r
}
