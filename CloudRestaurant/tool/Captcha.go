/*
	生成图片验证码
*/
package tool

import (
	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 生成图形化验证码
func GenerateCaptcha(context *gin.Context) {

	parameters := base64Captcha.ConfigCharacter{
		Height:             60,
		Width:              240,
		Mode:               3,
		ComplexOfNoiseText: 0,
		ComplexOfNoiseDot:  0,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		},
	}

	captchaId, captchaInterfaceInstance := base64Captcha.GenerateCaptcha("", parameters)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captchaInterfaceInstance)

	captchaResult := CaptchaResult{Id: captchaId, Base64Blob: base64blob}

	Success(context, map[string]interface{}{
		"captcha_result": captchaResult,
	})

}

// 验证输入的验证码是否和图片验证码相符
func VerifyCaptcha(id string, value string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(id, value)
	return verifyResult
}
