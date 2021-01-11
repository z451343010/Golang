/*
	Gin 框架
	表单绑定结构体
	ShouldJson
	使用 JSON 格式提交数据
*/
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

func main() {

	engine := gin.Default()

	// 解析前端提交的 JSON 格式的数据
	// http://localhost:8080/addStudent
	engine.POST("/addStudent", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		var person Person
		err := context.BindJSON(&person)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		context.Writer.Write([]byte("person：" +
			person.Name + " " + person.Sex + " " +
			strconv.Itoa(person.Age) + "\n"))

	})

	engine.Run()

}
