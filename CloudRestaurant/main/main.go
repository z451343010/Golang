/*
	GoWeb
	Gin框架
	在线订餐系统API
*/
package main

import (
	"fmt"
	"gin/CloudRestaurant/controller"
	"gin/CloudRestaurant/tool"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 路由设置
func registerRouter(router *gin.Engine) {

	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
	new(controller.FoodCategoryController).Router(router)
	new(controller.ShopController).Router(router)
	new(controller.GoodsController).Router(router)

}

// 跨域访问：cross origin resource share
func Cors() gin.HandlerFunc {

	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)

		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin,"+
				"access-control-allow-headers,%s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		// OPTIONS 嗅探
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		// 让程序继续处理请求
		context.Next()

	}

}

func main() {
	// 解析配置文件
	cfg, err := tool.ParseConfig("../config/app.json")

	if err != nil {
		panic(err.Error())
	}

	// 实例化数据库
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 初始化 redis 配置
	tool.InitRedisStore()

	app := gin.Default()

	// 设置全局跨域访问
	app.Use(Cors())

	// 设置 sessions
	tool.InitSession(app)

	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}
