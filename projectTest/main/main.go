package main

/*
	测试后端数据向前端发送请求
*/
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func main() {

	engine := gin.Default()

	// 设置 html 目录
	engine.LoadHTMLGlob("./html/*")

	// 加载静态资源如 css js
	// 文件夹 resources 需要和 main.go 所在的目录相同
	// 比如这里的 main.go 和 resources 都在 projectTest\main 目录下
	engine.Static("/resources", "./resources")

	// http://localhost:8080/helloHTML
	engine.GET("/helloHTML", func(context *gin.Context) {

		fullPath := context.FullPath()
		fmt.Println("请求的路径为：", fullPath)

		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": fullPath,
		})

	})

	// http://localhost:8080/test
	engine.POST("/test", func(context *gin.Context) {

		fullPath := context.FullPath()
		fmt.Println("请求的路径为：", fullPath)

		fmt.Println("test in")

		resp := &Response{Code: 1, Message: "OK", Data: fullPath}

		context.JSON(200, resp)

	})

	engine.Run()

}
