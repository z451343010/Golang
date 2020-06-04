package main

import (
	"fmt"
	"strconv"
)

// 使用系统内置的字符串转换函数
// 把字符串转换成整数
func main() {

	n, err := strconv.Atoi("123")

	if err != nil { // nil 相当于 null

		fmt.Println("转换失败，", err)

	} else {

		fmt.Println("转换成功，转换成：", n)

	}

	// 整数转换成字符串
	var str string = strconv.Itoa(123)
	fmt.Printf("str value = %v,str type = %T\n", str, str)

	var str3 = "abc"
	var str2 = "abc"
	fmt.Println(str3 == str2)

}
