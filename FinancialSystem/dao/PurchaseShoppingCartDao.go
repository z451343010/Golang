package dao

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type PurchaseShoppingCartDao struct {
	*tool.Orm
}

func (pscd *PurchaseShoppingCartDao) InsertPurchaseDetailToCart(psc model.PurchaseShoppingCart) int64 {

	result, err := pscd.Orm.InsertOne(&psc)

	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "PurchaseShoppingCartDao InsertPurchaseDetailToCart", err)
		return 0
	}

	return result

}

// 根据订单编号、商品ID 从购物车中删除记录
func (pscd *PurchaseShoppingCartDao) DeletePurchaseDetailFromCart(orderId int64, goodsId int64) int64 {

	purchaseShoppingCart := model.PurchaseShoppingCart{OrderId: orderId, GoodsId: goodsId}

	result, err := pscd.Delete(&purchaseShoppingCart)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "PurchaseShoppingCartDao DeletePurchaseDetailFromCart", err)
		return 0
	}

	return result

}

// 根据进货单编号查找商品详情
func (pscd *PurchaseShoppingCartDao) SelectPurchaseDetailByOrderId(orderId int64) ([]model.PurchaseShoppingCart, error) {

	var purchaseShoppingCarts []model.PurchaseShoppingCart

	err := pscd.Orm.Where("order_id = ?", orderId).Find(&purchaseShoppingCarts)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "PurchaseShoppingCartDao FindPurchaseDetailByOrderId", err)
		return nil, err
	}

	return purchaseShoppingCarts, nil

}
