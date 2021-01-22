package service

import (
	"gin/FinancialSystem/dao"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type UserService struct {
}

// 用户名 + 密码 登录
func (us *UserService) LoginPwd(username string, password string) *model.User {

	// 对前端输入的用户密码进行加密
	password = tool.EncoderSha256(password)

	userDao := dao.UserDao{tool.DbEngine}
	user := userDao.QueryByUsernameAndPwd(username, password)

	if user == nil {
		return nil
	}

	return user

}
