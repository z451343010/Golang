/*
	进货单详情
*/
package model

type PurchaseOrderDetail struct {
	OrderDetailId    int64   `xorm:"pk autoincr" json:"order_detail_id"` // 进货详情ID
	OrderId          int64   `xorm:"bigint" json:"order_id"`             // 进货单ID
	GoodsId          int64   `xorm:"int" json:"goods_id"`                // 商品ID
	GoodsUnitNum     int64   `xorm:"int" json:"goods_unit_num"`          // 进货单单位数量
	GoodsSingleNum   int64   `xorm:"int" json:"goods_single_num"`        // 进货单单件数量
	GoodsUnitPrice   float64 `xorm:"double" json:"goods_unit_price"`     // 进货单位进价
	GoodsSinglePrice float64 `xorm:"double" json:"goods_single_price"`   // 进货单单件进价
	OrderTotalPrice  float64 `xorm:"double" json:"order_total_price"`    // 该笔进货单总支出
	IsDebt           bool    `json:"is_debt"`                            // 该笔订单是否赊账
}
