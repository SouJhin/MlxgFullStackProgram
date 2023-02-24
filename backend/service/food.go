package service

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/define"
	"backend/models"
	"github.com/gin-gonic/gin"
)

var engine = models.Init()

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
	id := ctx.DefaultQuery("id", "-1")
	convertPrice, _ := strconv.Atoi(price)
	food := models.Food{
		Image:    "asdasd",
		Name:     name,
		ID:       id,
		Price:    convertPrice,
		Describe: "youde",
	}
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

func UpdateFood(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", define.DefaultPrice)
	price := ctx.DefaultQuery("price", define.DefaultPage)
	id := ctx.DefaultQuery("id", define.DefaultPage)
	convertPrice, _ := strconv.Atoi(price)
	food := models.Food{
		Image:    "asdasd",
		Name:     name,
		ID:       id,
		Price:    convertPrice,
		Describe: "youde",
	}
	fmt.Printf("food =====> 🚀🚀🚀 %v\n", food)
	value, err := engine.ID(id).Update(&food)
	fmt.Printf("value =====> 🚀🚀🚀 %v\n", value)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": food,
	})
}
