package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Log(ctx *gin.Context) {
	ctx.String(http.StatusOK, "service log")
}
