package service

import (
	"gin/FinancialSystem/dao"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type ShippingOrderDetailService struct {
}

// 插入出货单详情，生成出货单
func (sods *ShippingOrderDetailService) CommitShippingOrder(shippingOrderDetails []model.ShippingOrderDetail) {

	shippingOrderDetailDao := dao.ShippingOrderDetailDao{tool.DbEngine}
	shippingOrderDetailDao.CommitShippingDetail(shippingOrderDetails)

}
