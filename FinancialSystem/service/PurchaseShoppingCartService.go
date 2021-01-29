package service

import (
	"gin/FinancialSystem/dao"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type PurchaseShoppingCartService struct {
}

// 添加进货单详情到购物车
func (pscs *PurchaseShoppingCartService) AddPurchaseShopCart(psc model.PurchaseShoppingCart) bool {

	purchaseShoppingCartDao := dao.PurchaseShoppingCartDao{tool.DbEngine}
	result := purchaseShoppingCartDao.InsertPurchaseDetailToCart(psc)

	if result > 0 {
		return true
	}

	return false

}

// 根据 orderId goodsId 从购物车中删除某条记录
func (pscs *PurchaseShoppingCartService) RemovePurchaseShopCart(orderId int64, goodsId int64) bool {

	purchaseShoppingCartDao := dao.PurchaseShoppingCartDao{tool.DbEngine}

	result := purchaseShoppingCartDao.DeletePurchaseDetailFromCart(orderId, goodsId)

	if result > 0 {
		return true
	}

	return false

}

func (pscs *PurchaseShoppingCartService) FindPurchaseDetailByOrderId(orderId int64) []model.PurchaseShoppingCart {

	purchaseShoppingCartDao := dao.PurchaseShoppingCartDao{tool.DbEngine}
	purchaseShoppingCarts, err := purchaseShoppingCartDao.SelectPurchaseDetailByOrderId(orderId)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "PurchaseShoppingCartService FindPurchaseDetailByOrderId", err)
		return nil
	}

	return purchaseShoppingCarts

}
