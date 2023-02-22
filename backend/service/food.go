package service

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/define"
	"backend/models"
	"github.com/gin-gonic/gin"
)

// AddFood
// @Tags 添加食物
// @Summary 传入食物名称和价格
// @Param name query string false "请输入食物名称"
// @Param size query int false "请输入食物价格"
// Success 200 {string} json "{"code":"200", "msg", "data":""}"
// @Router /addFood [get]
func AddFood(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", define.DefaultPage)
	price := ctx.DefaultQuery("price", define.DefaultPrice)
	convertPrice, _ := strconv.Atoi(price)
	food := models.Food{
		Image:    "asdasd",
		Name:     name,
		ID:       "sssss",
		Price:    convertPrice,
		Describe: "youde",
	}
	engine := models.Init()
	_, err := engine.Insert(&food)
	findFood := make([]models.Food, 0)
	err = engine.Find(&findFood)
	for _, value := range findFood {
		fmt.Printf("food ====> %v\n", value)
	}
	fmt.Printf("summary %v\n\n", findFood)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"data":  findFood,
			"name":  name,
			"price": convertPrice,
		},
	})
}

// DeleteFood
// @Tags 删除食物
// @Summary 传入吧
// @Param size query int false "页数"
// @Param page query int false "页码"
// Success 200 {string} json "{"code":"200", "msg", "data":""}"
// @Router /deleteFood [post]
func DeleteFood() {
	fmt.Printf("1 =====> 🚀🚀🚀 %v\n", 1)
	fmt.Printf("1 =====> 🚀🚀🚀 %v\n", 1)
}
