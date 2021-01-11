/*
	Gin 框架
	表单绑定结构体
	ShouldBlindQuery
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 创建映射的结构体
type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

func main() {

	engine := gin.Default()

	// GET
	// http://localhost:8080/hello?name=Davie&classes=软件工程
	engine.GET("/hello", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		var student Student
		err := context.ShouldBindQuery(&student)

		if err != nil {
			return
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)

		context.Writer.Write([]byte("hello " +
			student.Name + " " + student.Classes + "\n"))

	})

	engine.Run()

}
