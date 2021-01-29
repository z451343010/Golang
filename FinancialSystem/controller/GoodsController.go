package controller

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/service"
	"gin/FinancialSystem/tool"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
}

func (gc *GoodsController) Router(engine *gin.Engine) {

	engine.POST("/user/add_goods", gc.addGoods)
	engine.POST("/user/count_unit_pirce", gc.countUintPrice)
	engine.GET("/user/find_all_goods", gc.findAllGoods)
	engine.GET("/user/find_goods_by_id", gc.findGoodsById)

}

// 添加商品
func (gc *GoodsController) addGoods(context *gin.Context) {

	var goodsParam model.Goods
	// 解析请求的参数
	err := tool.Decode(context.Request.Body, &goodsParam)
	if err != nil {
		tool.Failed(context, "插入商品，参数解析失败")
		return
	}

	goodsService := service.GoodsService{}
	result := goodsService.AddGoods(goodsParam)
	if result == 0 {
		tool.Failed(context, "插入商品失败")
		return
	}

	tool.Success(context, goodsParam)

}

// 计算商品单位价格
func (gc *GoodsController) countUintPrice(context *gin.Context) {

	var goodsParam model.Goods

	// 解析请求的参数
	err := tool.Decode(context.Request.Body, &goodsParam)
	if err != nil {
		tool.Failed(context, "计算商品单位价格，参数解析失败")
		return
	}

	// 计算商品单位进货价
	if goodsParam.GoodsUnitNum != 0 && goodsParam.GoodsPrice != 0 {
		goodsParam.GoodsUnitPrice = goodsParam.GoodsPrice * float64(goodsParam.GoodsUnitNum)
	}

	// 计算商品单位售价
	if goodsParam.GoodsUnitNum != 0 && goodsParam.SellingPrice != 0 {
		goodsParam.UnitSellingPrice = goodsParam.SellingPrice * float64(goodsParam.GoodsUnitNum)
	}

	tool.Success(context, goodsParam)

}

// 查找所有的商品
func (gc *GoodsController) findAllGoods(context *gin.Context) {

	var goods []model.Goods
	goodsService := service.GoodsService{}

	goods, err := goodsService.FindAllGoods()

	if err != nil {
		tool.Failed(context, "查找所有的商品发生错误")
		return
	}

	tool.Success(context, goods)

}

// 通过 ID 来查找商品
func (gc *GoodsController) findGoodsById(context *gin.Context) {

	// 获取前端传递过来的商品ID
	var goodsParam model.Goods
	err := context.BindJSON(&goodsParam)
	if err != nil {
		tool.Failed(context, "解析前端传递过来的商品ID失败")
		return
	}

	goodsId := goodsParam.GoodsId

	goodsService := service.GoodsService{}
	goods := goodsService.FindGoodsById(goodsId)

	tool.Success(context, goods)

}
