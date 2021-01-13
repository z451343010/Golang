/*
	GoWeb
	Gin框架
	中间件 middleWare
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义中间件函数
// 中间件函数满足的标准：
// 返回值类型为 HandlerFunc 的 func 函数
func RequestInfos() gin.HandlerFunc {

	return func(context *gin.Context) {

		fullPath := context.FullPath()
		// 获取请求的http类型
		method := context.Request.Method
		fmt.Println("请求路径：", fullPath)
		fmt.Println("请求类型：", method)

		// context.Next,将代码分为两个部分
		// 一部分请求之前执行，另一部分请求之后执行
		context.Next()

		fmt.Println("状态码：", context.Writer.Status())

	}

}

func main() {

	engine := gin.Default()

	// 使用中间件
	// engine.Use(RequestInfos())

	engine.GET("/query", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})

	// 单独让某个请求使用中间件 RequestInfos()
	// 加入参数 RequestInfos()
	engine.GET("/query2", RequestInfos(), func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})

	engine.Run()

}
