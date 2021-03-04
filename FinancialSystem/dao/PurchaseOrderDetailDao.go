package dao

import (
	"fmt"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
	"time"
)

type PurchaseOrderDetailDao struct {
	*tool.Orm
}

func (pod *PurchaseOrderDetailDao) CommitPurchaseDetail(purchaseOrderDetails []model.PurchaseOrderDetail) {

	session := pod.NewSession()
	defer session.Close()

	// 开启事务
	err := session.Begin()
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "PurchaseOrderDetailDao commitPurchaseDetail", err)
		return
	}

	var purchaseOrderTotalPrice float64

	// 把进货单详情插入数据库
	for _, value := range purchaseOrderDetails {

		purchaseOrderTotalPrice += float64(value.GoodsSingleNum)*value.GoodsSinglePrice +
			float64(value.GoodsUnitNum)*value.GoodsUnitPrice

		// 插入进货单详情
		_, err = session.InsertOne(&value)
		if err != nil {
			session.Rollback()
		}

		// 根据进货单详情，添加库存
		var inStock model.InStock
		inStock.GoodsId = value.GoodsId

		// 查看库存表中是否已经有该商品
		isInStock, err := session.Where("goods_id = ?", inStock.GoodsId).Get(&inStock)

		// 如果存在，则在原来的数量上累加
		if isInStock {
			inStock.GoodsSingleNum += value.GoodsSingleNum
			inStock.GoodsUnitNum += value.GoodsUnitNum
			_, err = session.Where("goods_id = ?", inStock.GoodsId).Update(&inStock)
			if err != nil {
				fmt.Println("update err = ", err)
				session.Rollback()
			}
		} else {
			// 如果不存在，则插入一条新的库存记录
			inStock.GoodsSingleNum += value.GoodsSingleNum
			inStock.GoodsUnitNum += value.GoodsUnitNum
			_, err = session.InsertOne(&inStock)
			if err != nil {
				session.Rollback()
			}
		}

	}

	// 插入进货单记录
	purchaseOrder := model.PurchaseOrder{}
	purchaseOrder.OrderId = purchaseOrderDetails[0].OrderId
	purchaseOrder.OrderMan = "高赟晖"
	timeNow := time.Now()
	timeString := timeNow.Format("2006-01-02 15:04:05")
	purchaseOrder.OrderDate = timeString
	purchaseOrder.OrderTotalPrice = purchaseOrderTotalPrice
	_, err = session.InsertOne(&purchaseOrder)
	if err != nil {
		session.Rollback()
	}

	// 提交事务
	err = session.Commit()
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "PurchaseOrderDetailDao commitPurchaseDetail", err)
		return
	}

}
