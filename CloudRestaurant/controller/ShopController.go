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
	engine.GET("/api/search_shops", sc.SearchShops)

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

	if len(shops) == 0 {
		tool.Failed(context, "暂未获取到商户信息")
		return
	}

	// 进一步获取店铺提供的服务
	for index, shop := range shops {

		shopServices := shopService.GetService(shop.Id)

		if len(shopServices) == 0 {
			shops[index].Supports = nil
		} else {
			shops[index].Supports = shopServices
		}
	}

	tool.Success(context, shops)

}

// 关键词搜索店铺信息
func (sc *ShopController) SearchShops(context *gin.Context) {

	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	keyword := context.Query("keyword")

	if keyword == "" {
		tool.Failed(context, "重新输入商铺名称")
		return
	}

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "118.6695"
		latitude = "24.8709"
	}

	// 执行真实的搜索逻辑
	shopService := service.ShopService{}
	shops := shopService.SearchShops(longitude, latitude, keyword)
	tool.Success(context, shops)

}
