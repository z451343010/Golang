package dao

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type UserDao struct {
	*tool.Orm
}

func (ud *UserDao) QueryByUsernameAndPwd(username string, password string) *model.User {

	var user model.User

	_, err := ud.Where("username = ? and password = ?", username, password).Get(&user)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "UserDao QueryByUsernameAndPwd", err)
		return nil
	}

	return &user

}
