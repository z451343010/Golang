package param

// 手机号 + 验证码登陆时，参数的传递
type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
