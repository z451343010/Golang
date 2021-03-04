package controller

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/service"
	"gin/FinancialSystem/tool"

	"github.com/gin-gonic/gin"
)

type ShippingOrderDetailController struct {
}

func (sodc *ShippingOrderDetailController) Router(engine *gin.Engine) {
	engine.POST("/user/commit_shipping_order", sodc.commitShippingOrder)
}

// 提交出货单 commitShippingOrder
func (sodc *ShippingOrderDetailController) commitShippingOrder(context *gin.Context) {

	// 接收前端传递过来的购物车参数
	var cartParam model.ShippingShoppingCart
	err := tool.Decode(context.Request.Body, &cartParam)
	if err != nil {
		tool.Failed(context, "出货单购物车详情，参数解析失败")
		return
	}

	// 获取出货单ID
	shippingShoppingCartService := service.ShippingShoppingCartService{}
	shippingShoppingCarts := shippingShoppingCartService.FindShippingDetailByOrderId(cartParam.OrderId)

	shippingDetailOrderId := tool.CreateOrderId()
	var shippingOrderDetails []model.ShippingOrderDetail

	for _, value := range shippingShoppingCarts {
		shippingOrderDetail := model.ShippingOrderDetail{OrderId: shippingDetailOrderId}
		shippingOrderDetail.GoodsId = value.GoodsId
		shippingOrderDetail.ShippingAddress = value.ShippingAddress
		shippingOrderDetail.Customer = value.Customer
		shippingOrderDetail.UnitPurchasePrice = value.UnitPurchasePrice
		shippingOrderDetail.UnitShippingPrice = value.UnitShippingPrice
		shippingOrderDetail.UnitNum = value.UnitNum
		shippingOrderDetail.SinglePurchasePrice = value.SinglePurchasePrice
		shippingOrderDetail.SingleShippingPrice = value.SingleShippingPrice
		shippingOrderDetail.ShippingUnitNum = value.ShippingUnitNum
		shippingOrderDetail.ShippingSingleNum = value.ShippingSingleNum
		shippingOrderDetail.IsDebt = value.IsDebt
		shippingOrderDetail.Sale = value.UnitShippingPrice*float64(value.ShippingUnitNum) + value.SingleShippingPrice*float64(value.ShippingSingleNum)
		shippingOrderDetail.Costing = value.UnitPurchasePrice*float64(value.ShippingUnitNum) + value.SinglePurchasePrice*float64(value.ShippingSingleNum)
		shippingOrderDetail.Profit = shippingOrderDetail.Sale - shippingOrderDetail.Costing
		shippingOrderDetail.ShippingMan = value.ShippingMan
		shippingOrderDetail.ChinesePrice = value.ChinesePrice
		shippingOrderDetails = append(shippingOrderDetails, shippingOrderDetail)
	}

	shippingOrderDetailService := service.ShippingOrderDetailService{}
	shippingOrderDetailService.CommitShippingOrder(shippingOrderDetails)

}
