/*
	Goweb
	Gin框架
	多数据格式返回请求结果
	JSON Format
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func main() {

	engine := gin.Default()

	// http://localhost:8080/helloJSON
	engine.GET("/helloJSON", func(context *gin.Context) {

		fullPath := context.FullPath()
		fmt.Println("请求的路径为：", fullPath)

		context.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "OK",
			"data":    fullPath,
		})

	})

	engine.GET("/jsonStruct", func(context *gin.Context) {

		fullPath := context.FullPath()
		fmt.Println("请求的路径为：", fullPath)

		resp := &Response{Code: 1, Message: "OK", Data: fullPath}

		context.JSON(200, resp)

	})

	engine.Run()

}
