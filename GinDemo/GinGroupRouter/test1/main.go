/*
	GoWeb
	Gin框架
	分组路由 group router
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	routerGroup := engine.Group("/user")

	routerGroup.POST("/register", func(context *gin.Context) {

		fullPath := context.FullPath()
		fmt.Println(fullPath)

		context.Writer.WriteString(fullPath)

	})

	routerGroup.POST("/login", func(context *gin.Context) {

		fullPath := context.FullPath()
		fmt.Println(fullPath)

		context.Writer.WriteString(fullPath)

	})

	engine.Run()

}
