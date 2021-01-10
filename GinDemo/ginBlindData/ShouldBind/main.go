/*
	Gin 框架
	表单绑定结构体
	ShouldBlind
*/
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Register struct {
	Username string `form:"username"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

func main() {

	engine := gin.Default()

	// http://localhost:8080/register
	engine.POST("/register", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		var register Register
		// ShouldBind
		err := context.ShouldBind(&register)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(register.Username)
		fmt.Println(register.Phone)
		fmt.Println(register.Password)

		context.Writer.Write([]byte("register：" +
			register.Username + " " + register.Phone + " " +
			register.Password + "\n"))

	})

	// 启动引擎
	engine.Run()

}
