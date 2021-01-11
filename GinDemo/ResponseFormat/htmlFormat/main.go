/*
	GoWeb
	Gin框架
	直接加载 html 模板
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	// 设置 html 目录
	engine.LoadHTMLGlob("./html/*")

	// 加载静态资源
	// 图片资源
	engine.Static("/img", "./img")

	// http://localhost:8080/helloHTML
	engine.GET("/helloHTML", func(context *gin.Context) {

		fullPath := context.FullPath()
		fmt.Println("请求的路径为：", fullPath)

		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": fullPath,
		})

	})

	engine.Run()

}
