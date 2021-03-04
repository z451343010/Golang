package service

import (
	"gin/FinancialSystem/dao"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type ShippingShoppingCartService struct {
}

// 添加出货单详情到购物车
func (sscs *ShippingShoppingCartService) AddShippingShopCart(ssc model.ShippingShoppingCart) bool {

	shippingShoppingCartDao := dao.ShippingShoppingCartDao{tool.DbEngine}
	result := shippingShoppingCartDao.InsertShippingDetailToCart(ssc)

	if result > 0 {
		return true
	}

	return false

}

// 根据 orderId goodsId 从出货单购物车中删除某条记录
func (sscs *ShippingShoppingCartService) RemoveShippingShopCart(orderId int64, goodsId int64) bool {

	shippingShoppingCartDao := dao.ShippingShoppingCartDao{tool.DbEngine}

	result := shippingShoppingCartDao.DeleteShippingDetailFromCart(orderId, goodsId)

	if result > 0 {
		return true
	}

	return false

}

// 根据出货单ID,查找出货单详情
func (sscs *ShippingShoppingCartService) FindShippingDetailByOrderId(orderId int64) []model.ShippingShoppingCart {

	shippingShoppingCartDao := dao.ShippingShoppingCartDao{tool.DbEngine}
	shippingShoppingCarts, err := shippingShoppingCartDao.SelectShippingDetailByOrderId(orderId)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "ShippingShoppingCartService FindShippingDetailByOrderId", err)
		return nil
	}

	return shippingShoppingCarts

}
