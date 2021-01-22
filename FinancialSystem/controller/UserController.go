package controller

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/service"
	"gin/FinancialSystem/tool"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) Router(engine *gin.Engine) {
	engine.POST("/user/user_login", uc.userLogin)
}

// 用户登录
// 用户名 + 密码 登录
func (uc *UserController) userLogin(context *gin.Context) {

	// 解析用户登录时传递的参数
	var userParam model.User

	// 解析请求的参数
	err := tool.Decode(context.Request.Body, &userParam)
	if err != nil {
		tool.Failed(context, "用户名 + 密码 参数解析失败")
		return
	}

	// 登录
	userService := service.UserService{}
	user := userService.LoginPwd(userParam.Username, userParam.Password)
	if user.UserId != 0 {
		tool.Success(context, user)
		return
	}

	tool.Failed(context, "登录失败,请重新校验用户名和密码")

}
