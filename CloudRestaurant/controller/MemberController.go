package controller

import (
	"encoding/json"
	"fmt"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/param"
	"gin/CloudRestaurant/service"
	"gin/CloudRestaurant/tool"
	"strconv"
	"time"

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
	engine.POST("/api/login_pwd", memberController.loginPwd)
	// 头像上传
	engine.POST("/api/upload/avator", memberController.uploadAvator)
	// 用户信息查询
	engine.GET("/api/userinfo", memberController.userInfo)

}

// 用户信息查询
func (memberController *MemberController) userInfo(context *gin.Context) {

}

// 用户名 + 密码 登录
func (memberController *MemberController) loginPwd(context *gin.Context) {

	// 1.解析用户名登陆传递参数
	var loginParam param.LoginParam
	// 解析请求的参数
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	// 2.验证验证码
	// validate := tool.VerifyCaptcha(loginParam.Id, loginParam.Value)
	// if !validate {
	// 	tool.Failed(context, "验证码不正确，请重新验证")
	// 	return
	// }

	// 3.登录
	memberService := service.MemberService{}
	member := memberService.LoginPwd(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		// 将用户信息保存到 session 中
		sess, _ := json.Marshal(member)
		err = tool.SetSession(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "保存 session 失败")
			return
		}

		tool.Success(context, &member)
		return
	}

	tool.Failed(context, "登录失败")

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

		sess, _ := json.Marshal(member)
		err = tool.SetSession(context, "user_id"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "保存 session 失败")
			return
		}

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
func (memberController *MemberController) verifyCaptcha(context *gin.Context) {

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

// 用户头像上传
func (memberController *MemberController) uploadAvator(context *gin.Context) {

	// 1.解析上传的参数 file、user_id
	userId := context.PostForm("user_id") // 用户id
	file, err := context.FormFile("avatar")
	if err != nil || userId == "" {
		tool.Failed(context, "参数解析失败")
		return
	}

	// 2.判断 user_id 对应的用户是否已经登录
	sess := tool.GetSession(context, "user_"+userId)
	if sess == nil {
		tool.Failed(context, "参数不合法")
		return
	}
	var member model.Member
	json.Unmarshal(sess.([]byte), &member)

	// 3.file 保存到本地
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(context, "头像更新失败")
		return
	}

	// 4.将保存后的文件本地路径，保存到用户表中的头像字段
	memberService := service.MemberService{}
	path := memberService.UploadAvatar(member.Id, fileName[1:])
	if path != "" {
		tool.Success(context, "http://localhost:8090"+path)
		return
	}

	// 5.返回结果
	tool.Failed(context, "用户头像上传失败")

}
