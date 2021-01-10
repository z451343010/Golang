package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// 创建一个engine
	engine := gin.Default()

	//
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		// 通过 FullPath
		// 能够拿到具体解析的是哪一个接口
		path := context.FullPath()

		// 打印请求的接口
		fmt.Println(path)

		// 获取链接?后面的参数值
		// 如果 name 获取不到，则默认值获取 hello
		name := context.DefaultQuery("name", "hello")
		fmt.Println(name)

		// 输出
		// Write 表示向前端（浏览器端）写入一个返回的数据
		context.Writer.Write([]byte("Hello Gin " + name + "\n"))

	})

	// 处理一个 POST 请求
	// http://localhost:8080/hello?name=Davie
	engine.Handle("POST", "/login", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		username := context.PostForm("username")
		phone := context.PostForm("phone")

		fmt.Println(username)
		fmt.Println(phone)

		context.Writer.Write([]byte(username + "登陆" + "\n"))

	})

	// 让 engine 运行起来
	engine.Run()

}
