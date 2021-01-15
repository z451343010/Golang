/*
	GoWeb
	Gin框架
	在线订餐系统API
*/
package main

import (
	"gin/CloudRestaurant/controller"
	"gin/CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

// 路由设置
func registerRouter(router *gin.Engine) {

	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)

}

func main() {
	// 解析配置文件
	cfg, err := tool.ParseConfig("../config/app.json")

	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()

	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}
