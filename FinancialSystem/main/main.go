/*
	进出货管理软件
*/
package main

import (
	"gin/FinancialSystem/controller"
	"gin/FinancialSystem/tool"

	"github.com/gin-gonic/gin"
)

// 路由设置
func registerRouter(router *gin.Engine) {
	new(controller.UserController).Router(router)
	new(controller.GoodsController).Router(router)
	new(controller.PurchaseOrderController).Router(router)
	new(controller.PurchaseShoppingCartController).Router(router)
}

func main() {

	// 解析配置文件
	cfg, err := tool.ParseConfig("../config/app.json")

	if err != nil {
		tool.LogRecoder(cfg, "main", err)
		return
	}

	// 实例化数据库
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		tool.LogRecoder(cfg, "main", err)
		return
	}

	app := gin.Default()

	// 注册路由
	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}
