package controller

import (
	"gin/FinancialSystem/tool"

	"github.com/gin-gonic/gin"
)

type PurchaseOrderController struct {
}

func (poc *PurchaseOrderController) Router(engine *gin.Engine) {

	engine.POST("/user/create_order_id", poc.returnOrderId)

}

// 返回 进货单ID 到页面
func (poc *PurchaseOrderController) returnOrderId(context *gin.Context) {

	purchaseOrderId := tool.CreateOrderId()
	tool.Success(context, purchaseOrderId)

}
