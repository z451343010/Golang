package controller

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/service"
	"gin/FinancialSystem/tool"

	"github.com/gin-gonic/gin"
)

type ShippingShoppingCartController struct {
}

func (sscc *ShippingShoppingCartController) Router(engine *gin.Engine) {
	engine.POST("/user/add_shipping_shopping_cart", sscc.addShippingShoppingCart)
	engine.POST("/user/remove_shipping_shopping_cart", sscc.removeShippingShoppingCart)
	engine.POST("/user/find_shipping_shopping_cart_by_order_id", sscc.findShippingShoppingCartByOrderId)
}

// 添加出货单详情到购物车
func (sscc *ShippingShoppingCartController) addShippingShoppingCart(context *gin.Context) {

	// 接收前端传递过来的出货单购物车参数
	var cartParam model.ShippingShoppingCart

	err := tool.Decode(context.Request.Body, &cartParam)
	if err != nil {
		tool.Failed(context, "出货单购物车详情，解析失败")
		return
	}

	// 查看商品库存
	inStockService := service.InStockService{}
	inStock := inStockService.FindInStockByGoodsId(cartParam.GoodsId)
	if cartParam.ShippingUnitNum > inStock.GoodsUnitNum || cartParam.ShippingSingleNum > inStock.GoodsSingleNum {
		tool.Failed(context, "出货的商品数量大于库存数量，请重新输入商品出货数量")
		return
	}

	shippingShoppingCartService := service.ShippingShoppingCartService{}
	result := shippingShoppingCartService.AddShippingShopCart(cartParam)
	if result {
		tool.Success(context, cartParam)
		return
	}

	tool.Failed(context, "插入出货单详情到购物车失败")

}

// 从出货单购物车中删除订单详情
func (sscc *ShippingShoppingCartController) removeShippingShoppingCart(context *gin.Context) {

	// 接收前端传递过来的购物车参数
	var cartParam model.ShippingShoppingCart

	err := tool.Decode(context.Request.Body, &cartParam)
	if err != nil {
		tool.Failed(context, "出货单购物车详情，参数解析失败")
		return
	}

	shippingShoppingCartService := service.ShippingShoppingCartService{}
	result := shippingShoppingCartService.RemoveShippingShopCart(cartParam.OrderId, cartParam.GoodsId)
	if result {
		tool.Success(context, "从出货单购物车中删除订单详情成功")
		return
	}

	tool.Failed(context, "从出货单购物车中删除订单详情失败")

}

// 根据出货单编号查找购物车详情记录
func (sscc *ShippingShoppingCartController) findShippingShoppingCartByOrderId(context *gin.Context) {

	// 接收前端传递过来的购物车参数
	var cartParam model.ShippingShoppingCart
	err := tool.Decode(context.Request.Body, &cartParam)
	if err != nil {
		tool.Failed(context, "购物车详情，参数解析失败")
		return
	}

	shippingShoppingCartService := service.ShippingShoppingCartService{}

	shippingShoppingCarts := shippingShoppingCartService.FindShippingDetailByOrderId(cartParam.OrderId)
	if shippingShoppingCarts != nil {
		tool.Success(context, shippingShoppingCarts)
		return
	}

	tool.Failed(context, "根据出货单ID查询购物车是失败")

}
