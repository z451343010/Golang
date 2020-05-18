package main

import "fmt"

// 实现登录验证，有三次机会，如果用户名为“张无忌”
// 密码为“888”，提示登录成功
// 否则提示还有几次机会
func main() {

	var username string
	var password string
	var opNum int = 3

	for {

		fmt.Println("请输入用户名：")
		fmt.Scanln(&username)
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)

		if username == "张无忌" && password == "888" {

			fmt.Println("登录成功")
			break

		} else {

			opNum--

			if opNum > 0 {
				fmt.Println("用户名密码错误，还有", opNum, "次机会")
				fmt.Println()
			}

			if opNum == 0 {
				fmt.Println("您的密码出错次数超过了安全限制，已经冻结了您的账户")
				break
			}

		}

	}

}
