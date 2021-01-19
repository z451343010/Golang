package controller

import (
	"gin/CloudRestaurant/service"
	"gin/CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
}

func (sc *ShopController) Router(engine *gin.Engine) {

	engine.GET("/api/shops", sc.GetShopList)

}

// 获取商铺列表
func (sc *ShopController) GetShopList(context *gin.Context) {

	// 获取经度
	longitude := context.Query("longitude")
	// 获取纬度
	latitude := context.Query("latitude")

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "118.6695"
		latitude = "24.8709"
	}

	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)

	if len(shops) != 0 {
		tool.Success(context, shops)
		return
	}

	tool.Failed(context, "暂未获取到商户信息")

}
