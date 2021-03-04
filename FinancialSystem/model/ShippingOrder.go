/*
	出货单
*/
package model

type ShippingOrder struct {
	OrderId           int64   `xorm:"pk not null" json:"order_id"`       // 出货单ID
	OrderTotalCosting float64 `xorm:"double" json:"order_total_costing"` // 出货单总成本
	OrderTotalSale    float64 `xorm:"double" json:"order_total_sale"`    // 出货单总销售额
	OrderTotalProfit  float64 `xorm:"double" json:"order_total_profit"`  // 出货单总利润
	OrderMan          string  `xorm:"varchar(12)" json:"order_man"`      // 出货人
	OrderDate         string  `xorm:"varchar(25)" json:"order_date"`     // 出货日期
	IsDebt            bool    `json:"is_debt"`                           // 是否赊账
}
