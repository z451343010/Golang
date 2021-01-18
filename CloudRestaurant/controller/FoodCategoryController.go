package controller

import (
	"gin/CloudRestaurant/service"
	"gin/CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {

	engine.GET("/api/food_category", fcc.foodCategory)

}

func (fcc *FoodCategoryController) foodCategory(context *gin.Context) {

	foodCategoryService := &service.FoodCategoryService{}

	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(context, "食品种类获取失败")
		return
	}

	// 转换格式
	// 拼接图片路径
	// for _,category := range categories {

	// 	// 拼接图片 url
	// 	if category.ImageUrl != "" {
	// 		category.ImageUrl = tool.FileServerAddr() + "/" + category.ImageUrl
	// 	}

	// }

	tool.Success(context, categories)

}
