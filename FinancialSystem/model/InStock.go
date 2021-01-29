/*
	商品库存
*/
package model

type InStock struct {
	GoodsId        int64 `xorm:"pk not null" json:"goods_id"` // 商品ID
	GoodsUnitNum   int64 `xorm:"int" json:"goods_unit_num"`   // 商品单位数量
	GoodsSingleNum int64 `xorm:"int" json:"goods_single_num"` // 商品单件数量
}
