package service

import (
	"gin/FinancialSystem/dao"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type PurchaseOrderDetailService struct {
}

// 插入进货单详情，生成进货单
func (pods *PurchaseOrderDetailService) CommitPurchaseOrder(purchaseOrderDetals []model.PurchaseOrderDetail) {

	purchaseOrderDetailDao := dao.PurchaseOrderDetailDao{tool.DbEngine}
	purchaseOrderDetailDao.CommitPurchaseDetail(purchaseOrderDetals)

}
