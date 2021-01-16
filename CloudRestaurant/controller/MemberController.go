package controller

import (
	"fmt"
	"gin/CloudRestaurant/param"
	"gin/CloudRestaurant/service"
	"gin/CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

// 解析 /sendMessage
func (memberController *MemberController) Router(engine *gin.Engine) {

	engine.GET("/sendMessage", memberController.sendMessage)
	engine.GET("/api/sendcode", memberController.sendSmsCode)
	engine.POST("/api/login_sms", memberController.smsLogin)
	engine.GET("/api/captcha", memberController.captcha)
	engine.POST("/api/verifycaptcha", memberController.verifyCaptcha)

}

// 发送短信验证码
// http://localhost:8090/api/sendcode?phone=13599263379
func (memberController *MemberController) sendSmsCode(context *gin.Context) {

	// 发送验证码
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}

	// 调用阿里云的短信发送验证功能
	ms := service.MemberService{}
	isSend := ms.SendSmsCode(phone)
	if isSend {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "短信发送成功",
		})
		return
	}

	context.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "短信发送失败",
	})

}

func (memberController *MemberController) sendMessage(context *gin.Context) {

	memberService := service.MemberService{}

	phone := context.DefaultQuery("phone", "")
	memberService.SendCode(phone)

	context.JSON(200, map[string]interface{}{
		"message": "sendMessage cloudrestaurant",
	})
}

// 手机号 + 短信登陆
func (memberController *MemberController) smsLogin(context *gin.Context) {

	var smsLoginParam param.SmsLoginParam
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		fmt.Println("====== SmsLoginParam ======")
		fmt.Println("解析参数出错")
		tool.Failed(context, "参数解析失败")
		return
	}

	// 完成手机验证 + 验证码登录
	service := service.MemberService{}
	member := service.SmsLogin(smsLoginParam)

	// 登录成功
	if member != nil {
		tool.Success(context, member)
		return
	}

	tool.Failed(context, "登录失败")

}

// 生成图片验证码
func (memberController *MemberController) captcha(context *gin.Context) {

	// 生成图形化验证码，并返回客户端
	tool.GenerateCaptcha(context)

}

// 验证图片验证码是否正确
func (MemberController *MemberController) verifyCaptcha(context *gin.Context) {

	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	result := tool.VerifyCaptcha(captcha.Id, captcha.VerifyValue)
	if result {
		fmt.Println("图形验证码验证通过")
	} else {
		fmt.Println("图形验证码验证失败")
	}
}
