/*
	出货单详情购物车
*/
package model

type ShippingShoppingCart struct {
	ShoppingCartId      int64   `xorm:"pk autoincr" json:"shopping_cart_id"`  // 购物车ID
	OrderId             int64   `xorm:"bigint" json:"order_id"`               // 出货单ID
	GoodsId             int64   `xorm:"int" json:"goods_id"`                  // 商品ID
	ShippingAddress     string  `xorm:"varchar(255)" json:"shipping_address"` // 出货地
	Customer            string  `xorm:"varchar(20)" json:"customer"`          // 客户名称
	UnitPurchasePrice   float64 `xorm:"double" json:"unit_purchase_price"`    // 单位进货价
	UnitShippingPrice   float64 `xorm:"double" json:"unit_shipping_price"`    // 单位出货价
	UnitNum             int64   `xorm:"int" json:"unit_num"`                  // 单位数量
	SinglePurchasePrice float64 `xorm:"double" json:"single_purchase_price"`  // 单件进货价
	SingleShippingPrice float64 `xorm:"double" json:"single_shipping_price"`  // 单件出货价
	ShippingUnitNum     int64   `xorm:"int" json:"shipping_unit_num"`         // 单位出货数量
	ShippingSingleNum   int64   `xorm:"int" json:"shipping_single_num"`       // 单件出货数量
	IsDebt              bool    `json:"is_debt"`                              // 是否赊账
	Sale                float64 `xorm:"double" json:"sale"`                   // 销售额
	Costing             float64 `xorm:"double" json:"costing"`                // 成本
	Profit              float64 `xorm:"double" json:"profit"`                 // 利润
	ShippingMan         string  `xorm:"varchar(20)" json:"shipping_man"`      // 出货人
	ChinesePrice        string  `xorm:"varchar(255)" json:"chinese_price"`    // 用中文显示的价格
}
