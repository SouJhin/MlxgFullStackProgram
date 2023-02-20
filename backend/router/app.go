package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.GET("/log", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"cao": "hahas",
		})
	})
}
