package dao

import (
	"fmt"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
	"time"
)

type ShippingOrderDetailDao struct {
	*tool.Orm
}

func (sod *ShippingOrderDetailDao) CommitShippingDetail(shippingOrderDetails []model.ShippingOrderDetail) {

	session := sod.NewSession()
	defer session.Close()

	// 开启事务
	err := session.Begin()
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "ShippingOrderDetailDao CommitShippingDetail", err)
		return
	}

	var orderTotalCosting float64
	var orderTotalSale float64
	var orderTotalProfit float64

	// 把出货单详情插入数据库
	for _, value := range shippingOrderDetails {

		// 计算该订单总成本、总销售额、总利润
		orderTotalCosting += value.Costing
		orderTotalSale += value.Sale
		orderTotalProfit += value.Profit

		// 插入出货单详情
		_, err = session.InsertOne(&value)

		if err != nil {
			session.Rollback()
		}

		// 根据进货单详情，更新库存
		inStock := model.InStock{GoodsId: value.GoodsId}
		isInStock, err := session.Where("goods_id = ?", inStock.GoodsId).Get(&inStock)

		if isInStock {
			inStock.GoodsSingleNum -= value.ShippingSingleNum
			inStock.GoodsUnitNum -= value.ShippingUnitNum
			_, err = session.Where("goods_id = ?", inStock.GoodsId).Update(&inStock)

			if err != nil {
				fmt.Println("update err = ", err)
				session.Rollback()
			}
		}

	}

	// 插入出货单记录
	shippingOrder := model.ShippingOrder{}
	shippingOrder.OrderId = shippingOrderDetails[0].OrderId
	shippingOrder.OrderMan = "高赟晖"
	timeNow := time.Now()
	timeString := timeNow.Format("2006年01月02日 15:04:05")
	shippingOrder.OrderDate = timeString
	shippingOrder.OrderTotalCosting = orderTotalCosting
	shippingOrder.OrderTotalSale = orderTotalSale
	shippingOrder.OrderTotalProfit = orderTotalProfit
	_, err = session.InsertOne(&shippingOrder)

	if err != nil {
		session.Rollback()
	}

	// 提交事务
	err = session.Commit()
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "ShippingOrderDetailDao CommitShippingDetail", err)
		return
	}

}
