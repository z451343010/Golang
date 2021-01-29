/*
	进货单
*/
package model

type PurchaseOrder struct {
	OrderId         int64   `xorm:"pk not null" json:"order_id"`     // 进货单ID
	OrderTotalPrice float64 `xorm:"double" json:"order_total_price"` // 进货单总支出
	OrderMan        string  `xorm:"varchar(12)" json:"order_man"`    // 进货人
	OrderDate       string  `xorm:"varchar(25)" json:"order_date"`   // 进货日期
	IsDebt          bool    `json:"is_debt"`                         // 是否赊账
}
