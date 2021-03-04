package controller

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/service"
	"gin/FinancialSystem/tool"

	"github.com/gin-gonic/gin"
)

type PurchaseOrderDetailController struct {
}

func (podc *PurchaseOrderDetailController) Router(engine *gin.Engine) {
	engine.POST("/user/commit_purchase_order", podc.commitPurchaseOrder)
}

// 提交订单
func (podc *PurchaseOrderDetailController) commitPurchaseOrder(context *gin.Context) {

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

	purchaseDetailOrderId := tool.CreateOrderId()
	var purchaseOrderDetails []model.PurchaseOrderDetail

	for _, value := range purchaseShoppingCarts {
		purchaseOrderDetail := model.PurchaseOrderDetail{OrderId: purchaseDetailOrderId}
		purchaseOrderDetail.GoodsId = value.GoodsId
		purchaseOrderDetail.GoodsUnitNum = value.GoodsUnitNum
		purchaseOrderDetail.GoodsSingleNum = value.GoodsSingleNum
		purchaseOrderDetail.GoodsUnitPrice = value.GoodsUnitPrice
		purchaseOrderDetail.GoodsSinglePrice = value.GoodsSinglePrice
		purchaseOrderDetail.OrderTotalPrice = value.OrderTotalPrice
		purchaseOrderDetail.IsDebt = value.IsDebt
		purchaseOrderDetails = append(purchaseOrderDetails, purchaseOrderDetail)
	}

	purchaseOrderDetailService := service.PurchaseOrderDetailService{}
	purchaseOrderDetailService.CommitPurchaseOrder(purchaseOrderDetails)

}
