/*
	单独分离 GET 请求
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	// http://localhost:8080/hello?name=Davie
	engine.GET("/hello", func(context *gin.Context) {

		name := context.Query("name")
		fmt.Println(name)

		context.Writer.Write([]byte("hello " + name + "\n"))

	})

	// POST 方法
	// http://localhost:8080/login
	engine.POST("/login", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		username, exist := context.GetPostForm("username")
		phone, exist := context.GetPostForm("phone")

		// exist 为 bool 类型
		// 代表字段是否存在
		if exist {
			context.Writer.Write([]byte("hello " + username +
				" " + phone + "\n"))
		}

	})

	// DELETE 方法
	engine.DELETE("/user/:id", func(context *gin.Context) {
		userID := context.Param("id")
		fmt.Println(userID)
		context.Writer.Write([]byte("delete 用户ID：" + userID))
	})

	engine.Run()

}
