package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	// 设置 html 目录
	engine.LoadHTMLGlob("../html/*")

	engine.GET("/helloHTML", func(context *gin.Context) {

		fullPath := context.FullPath()

		context.HTML(http.StatusOK, "test.html", gin.H{
			"fullPath": fullPath,
		})

	})

	engine.Run()

}
