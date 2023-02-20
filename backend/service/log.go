package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Log(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "fuck u",
	})
}
