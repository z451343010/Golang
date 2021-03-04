package dao

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type ShippingShoppingCartDao struct {
	*tool.Orm
}

func (pscd *ShippingShoppingCartDao) InsertShippingDetailToCart(ssc model.ShippingShoppingCart) int64 {

	result, err := pscd.Orm.InsertOne(&ssc)

	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "ShippingShoppingCartDao InsertShippingDetailToCart", err)
		return 0
	}

	return result

}

// 根据订单编号、商品ID 从出货单购物车中删除记录
func (sscd *ShippingShoppingCartDao) DeleteShippingDetailFromCart(orderId int64, goodsId int64) int64 {

	shippingShoppingCart := model.ShippingShoppingCart{OrderId: orderId, GoodsId: goodsId}

	result, err := sscd.Delete(&shippingShoppingCart)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "ShippingShoppingCartDao DeleteShippingDetailFromCart", err)
		return 0
	}

	return result

}

// 根据出货单编号查找商品详情
func (sscd *ShippingShoppingCartDao) SelectShippingDetailByOrderId(orderId int64) ([]model.ShippingShoppingCart, error) {

	var shippingShoppingCarts []model.ShippingShoppingCart

	err := sscd.Orm.Where("order_id = ?", orderId).Find(&shippingShoppingCarts)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "ShippingShoppingCartDao SelectShippingDetailByOrder", err)
		return nil, err
	}

	return shippingShoppingCarts, nil

}
