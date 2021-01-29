package controller

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/service"
	"gin/FinancialSystem/tool"

	"github.com/gin-gonic/gin"
)

type PurchaseShoppingCartController struct {
}

func (pscc *PurchaseShoppingCartController) Router(engine *gin.Engine) {
	engine.POST("/user/add_purchase_shopping_cart", pscc.addPurchaseShoppingCart)
	engine.POST("/user/remove_purchase_shopping_cart", pscc.removePurchaseShoppingCart)
	engine.POST("/user/find_purchase_shopping_cart_by_order_id", pscc.findPurchaseShoppingCartByOrderId)
}

// 添加进货单详情到购物车
func (pscc *PurchaseShoppingCartController) addPurchaseShoppingCart(context *gin.Context) {

	// 接收前端传递过来的购物车参数
	var cartParam model.PurchaseShoppingCart

	err := tool.Decode(context.Request.Body, &cartParam)
	if err != nil {
		tool.Failed(context, "购物车详情，参数解析失败")
		return
	}

	purchaseShoppingCartService := service.PurchaseShoppingCartService{}
	result := purchaseShoppingCartService.AddPurchaseShopCart(cartParam)
	if result {
		tool.Success(context, cartParam)
		return
	}

	tool.Failed(context, "插入进货单详情到购物车失败")

}

// 从购物车中删除订单详情
func (pscc *PurchaseShoppingCartController) removePurchaseShoppingCart(context *gin.Context) {

	// 接收前端传递过来的购物车参数
	var cartParam model.PurchaseShoppingCart

	err := tool.Decode(context.Request.Body, &cartParam)
	if err != nil {
		tool.Failed(context, "购物车详情，参数解析失败")
		return
	}

	// 获取订单ID
	orderId := cartParam.OrderId
	// 获取商品ID
	goodsId := cartParam.GoodsId

	purchaseShoppingCartService := service.PurchaseShoppingCartService{}
	result := purchaseShoppingCartService.RemovePurchaseShopCart(orderId, goodsId)
	if result {
		tool.Success(context, "从购物车中删除订单详情成功")
		return
	}

	tool.Failed(context, "从购物车中删除订单详情失败")

}

// 根据进货单编号查找购物车详情记录
func (pscc *PurchaseShoppingCartController) findPurchaseShoppingCartByOrderId(context *gin.Context) {

	// 接收前端传递过来的购物车参数
	var cartParam model.PurchaseShoppingCart
	err := tool.Decode(context.Request.Body, &cartParam)
	if err != nil {
		tool.Failed(context, "购物车详情，参数解析失败")
		return
	}

	// 获取订单ID
	orderId := cartParam.OrderId
	purchaseShoppingCartService := service.PurchaseShoppingCartService{}

	purchaseShoppingCarts := purchaseShoppingCartService.FindPurchaseDetailByOrderId(orderId)
	if purchaseShoppingCarts != nil {
		tool.Success(context, purchaseShoppingCarts)
		return
	}

	tool.Failed(context, "根据进货单ID查询购物车失败")

}
