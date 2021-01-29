/*
	商品结构体
*/
package model

type Goods struct {
	GoodsId          int64   `xorm:"pk autoincr" json:"goods_id"`            // 商品ID
	GoodsName        string  `xorm:"varchar(20)" json:"goods_name"`          // 商品名
	GoodsUnit        string  `xorm:"varchar(12) "json:"goods_unit"`          // 商品单位（如：箱）
	GoodsUnitPrice   float64 `xorm:"double" json:"goods_unit_price"`         // 商品单位进货价（如:每箱x元）
	UnitSellingPrice float64 `xorm:"double" json:"unit_selling_price"`       // 商品单位售价（如：每箱卖x元）
	GoodsSingleUnit  string  `xorm:"varchar(12)" json:"goods_single_unit"`   // 商品单件单位（如：瓶）
	GoodsPrice       float64 `xorm:"double" json:"goods_price"`              // 商品单件进货价
	SellingPrice     float64 `xorm:"double" json:"selling_price"`            // 商品单件售价
	GoodsUnitNum     int     `xorm:"int" json:"goods_unit_num"`              // 商品单位数量（如：每箱12瓶）
	PuchasingAddress string  `xorm:"varchar(255)" json:"purchasing_address"` // 商品进货地
}
