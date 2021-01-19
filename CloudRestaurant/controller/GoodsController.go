package controller

import (
	"gin/CloudRestaurant/service"
	"gin/CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
}

func (gc *GoodsController) Router(engine *gin.Engine) {
	engine.GET("/api/foods", gc.getGoods)
}

// 获取某个商户下面的商品
func (gc *GoodsController) getGoods(context *gin.Context) {

	shopId, exist := context.GetQuery("shop_id")

	if !exist {
		tool.Failed(context, "请求参数错误")
		return
	}

	goodsService := service.GoodsService{}
	goods := goodsService.GetFoods(shopId)

	tool.Success(context, goods)

}
