/*
	Goweb
	Gin框架
	多数据格式返回请求结果
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	// http://localhost:8080/helloByte
	engine.GET("/helloByte", func(context *gin.Context) {

		fullPath := context.FullPath()

		fmt.Println("请求的路径为：", fullPath)

		context.Writer.Write([]byte(fullPath + "\n"))

	})

	// 通过 string 的方法返回数据
	engine.GET("/helloString", func(context *gin.Context) {

		fullPath := context.FullPath()

		fmt.Println("请求的路径为：", fullPath)

		context.Writer.WriteString(fullPath)

	})

	engine.Run()

}
