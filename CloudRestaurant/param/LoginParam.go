/*
	用户名
	验证码
	密码
	参数传递
*/
package param

type LoginParam struct {
	Name     string `json:"name"`  // 用户名
	Password string `json:"pwd"`   // 密码
	Id       string `json:"id"`    // 验证码id
	Value    string `json:"value"` // 验证码输入值
}
