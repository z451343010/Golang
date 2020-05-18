package main

import "fmt"

// 判断用户名和密码
// 用户名是否为“张三”
// 密码是否为“1234”
// 如果是，提示登录成功，否则提示登录失败
func main() {

	var username string
	var password string
	fmt.Println("请输入您的用户名：")
	fmt.Scanln(&username)
	fmt.Println("请输入您的密码：")
	fmt.Scanln(&password)

	if username == "张三" && password == "1234" {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}

}
