package router

import (
	service "backend/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("log", service.Log)
	r.GET("newFood", service.AddFood)
	return r
}
