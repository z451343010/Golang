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

// 根据 用户名 + 密码 查询数据
func (memberDao *MemberDao) QueryByNameAndPwd(name string, password string) *model.Member {

	var member model.Member

	// 密码保存进入数据库时，对其加密
	password = tool.EncoderSha256(password)

	_, err := memberDao.Where("user_name = ? and password = ?", name, password).Get(&member)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &member

}

// 更新用户的头像地址字段
func (memberDao *MemberDao) UpdateMemberAvatar(userId int64, fileName string) int64 {

	member := model.Member{Avatar: fileName}
	result, err := memberDao.Where("id = ? ", userId).Update(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result

}

// 根据用户的ID进行查询
func (memberDao *MemberDao) QueryMemberById(id int64) *model.Member {

	var member model.Member
	_, err := memberDao.Orm.Where(" id = ? ", id).Get(&member)
	if err != nil {
		fmt.Println("根据 userId 查询失败")
		return nil
	}

	return &member

}
