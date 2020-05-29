package main

import (
	"fmt"
	"strings"
)

// 1）编写一个函数makeSuffix(suffix string)可以接受一个文件后缀名
// （比如.jpg），并返回一个闭包。
// 2）调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如.jpg）
// 则返回文件名.jpg，如果已经有.jpg后缀，则返回原文件名。
// 3）要求使用闭包的方式完成。
// 4）strings.HasSuffix
func makeSuffix(suffix string) func(string) string {

	return func(fileStr string) string {

		if strings.HasSuffix(fileStr, ".jpg") {
			return fileStr
		} else {
			return fileStr + suffix
		}

	}

}

func main() {

	var fileStr string
	fmt.Println("请输入文件名")
	fmt.Scanln(&fileStr)

	// 定义一个函数变量f
	// 用于接收闭包返回的函数,返回一个闭包
	suffixFunction := makeSuffix(".jpg")

	fmt.Println(suffixFunction(fileStr))
	fmt.Println(suffixFunction("333"))

}
