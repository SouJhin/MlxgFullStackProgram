package service

import (
	"net/http"
	"strconv"

	"backend/define"
	"backend/models"
	"github.com/gin-gonic/gin"
)

func AddFood(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", define.DefaultPage)
	price := ctx.DefaultQuery("price", define.DefaultPrice)
	convertPrice, _ := strconv.Atoi(price)
	models.InsetFood(name, convertPrice)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
