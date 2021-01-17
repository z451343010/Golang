package service

import (
	"encoding/json"
	"fmt"
	"gin/CloudRestaurant/dao"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/param"
	"gin/CloudRestaurant/tool"
	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type MemberService struct {
}

// 调用阿里云服务
// 实现发送验证码短信功能
func (ms *MemberService) SendSmsCode(phone string) bool {

	// 1.产生一个验证码
	code := fmt.Sprintf("%06v",
		rand.New(rand.NewSource(time.Now().UnixNano())).
			Int31n(1000000))

	// 2.调用阿里云的 sdk 完成发送
	smsConfig := tool.GetConfig().Sms

	client, err := dysmsapi.NewClientWithAccessKey(smsConfig.RegionId,
		smsConfig.AppKey, smsConfig.AppSecret)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = smsConfig.SignName
	request.TemplateCode = smsConfig.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(response)

	// 3.接受返回结果，并判断发送状态
	// 短信验证码发送成功
	if response.Code == "OK" {
		return true
	}

	return false

}

func (ms *MemberService) SendCode(phone string) bool {

	// 产生一个随机验证码
	code := fmt.Sprintf("%06v",
		rand.New(rand.NewSource(time.Now().UnixNano())).
			Int31n(1000000))

	/*
		由于没有购买阿里云的短信服务
		因此这里省略一个
		从阿里云发送短信验证码到手机的业务逻辑
	*/

	// 将验证码的内容保存到数据库当中
	smsCode := model.SmsCode{Phone: phone, Code: code,
		BizId: "1065599", CreateTime: time.Now().Unix()}

	// 执行 dao 层，插入数据到数据库
	memberDao := dao.MemberDao{tool.DbEngine}
	result := memberDao.InsertCode(smsCode)

	if result > 0 {
		return true
	}

	return false

}

// 用户 手机号+验证码 登录
func (memberService *MemberService) SmsLogin(loginParam param.SmsLoginParam) *model.Member {

	// 1.获取到手机号和验证码

	// 2.验证手机号+验证码是否正确
	md := dao.MemberDao{}
	sms := md.ValidateSmsCode(loginParam.Phone, loginParam.Code)
	if sms.Id == 0 {
		return nil
	}

	// 3.根据手机号 member 表中查询记录
	member := md.QueryByPhone(loginParam.Phone)
	if member.Id != 0 {
		return member
	}

	// 4.新创建一个 member 记录，并保存
	user := model.Member{}
	user.UserName = loginParam.Phone
	user.Mobile = loginParam.Phone
	user.RegisterTime = time.Now().Unix()
	user.Id = md.InsertMember(user)

	return &user

}

// 用户名 + 密码 登录
func (memberService *MemberService) LoginPwd(name string, password string) *model.Member {

	// 1.使用用户名 + 密码 查询用户信息
	memberDao := dao.MemberDao{tool.DbEngine}
	member := memberDao.QueryByNameAndPwd(name, password)
	if member.Id != 0 {
		return member
	}

	// 2.用户不存在，作为新用户插入数据
	user := model.Member{}
	user.UserName = name
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()

	result := memberDao.InsertMember(user)
	user.Id = result

	return &user

}
