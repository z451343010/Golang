package model

/*
	用户结构体
*/
type User struct {
	UserId   int64   `xorm:"pk autoincr" json:"user_id"`   // 用户ID
	Username string  `xorm:"varchar(20)" json:"username"`  // 用户名
	Password string  `xorm:"varchar(255)" json:"password"` // 密码
	Balance  float64 `xorm:"double" json:"balance"`        // 余额
	Name     string  `xorm:"varchar(20)" json:"name"`      // 真实姓名
}
