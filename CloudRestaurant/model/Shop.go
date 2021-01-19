/*
	商家
*/
package model

type Shop struct {
	Id                 int64   `xorm:"pk autoincr" json:"id"`             // 商家ID
	Name               string  `xorm:"varchar(12)" json:"name"`           // 商家名
	PromotionInfo      string  `xorm:"varchar(30)" json:"promotion_info"` // 宣传信息
	Address            string  `xorm:"varchar(100)" json:"address"`       // 地址
	Phone              string  `xorm:"varchar(11)" json:"phone"`          // 联系电话
	Status             int     `xorm:"tinyint" json:"status"`             // 店铺营业状态
	Longitude          float64 `xorm:"" json:"longitude"`                 // 经度
	Latitude           float64 `xorm:"" json:"latitude"`                  // 纬度
	ImagePath          string  `xorm:"varchar(255)" json:"image_path"`    // 图片路径
	Rating             float32 `xorm:"float" json:"rating"`               // 商铺评分
	RatingCount        int64   `xorm:"int" json:"rating_count"`           // 评分总数
	RecentOrderNum     int64   `xorm:"int" json:"recent_order_num"`       // 当前订单总数
	MinimumOrderAmount int32   `xorm:"int" json:"minimum_order_amount"`   // 配送费起价
	DeliveryFee        int32   `xorm:"int" json:"delivery_fee"`           // 配送费
	OpeningHours       string  `xorm:"varchar(20)" json:"opening_hours"`  // 营业时间

}
