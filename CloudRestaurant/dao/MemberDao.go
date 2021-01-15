package dao

import (
	"fmt"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/tool"
	"log"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {

	// 解析配置文件
	cfg, err := tool.ParseConfig("../config/app.json")
	if err != nil {
		panic(err.Error())
	}

	// 调用 tool.OrmEngine 建立数据库连接
	var test *tool.Orm
	test, err = tool.OrmEngine(cfg)

	result, err := test.InsertOne(&sms)

	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

// 验证手机号+验证码
func (memberDao *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {

	conn, err := tool.GetOrmConnect()

	var sms model.SmsCode
	_, err = conn.Where(" phone = ? and code = ?", phone,
		code).Get(&sms)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &sms

}

// 查询数据库中是否存在该手机号
func (memberDao *MemberDao) QueryByPhone(phone string) *model.Member {

	var member model.Member
	conn, err := tool.GetOrmConnect()

	_, err = conn.Where("mobile = ? ", phone).Get(&member)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &member

}

// 插入用户
func (memberDao *MemberDao) InsertMember(member model.Member) int64 {

	conn, err := tool.GetOrmConnect()

	result, err := conn.InsertOne(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return result

}
