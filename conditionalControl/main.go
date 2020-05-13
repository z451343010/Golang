package main

import "fmt"

// 如果年龄大于18，输出“你的年龄大于18，要对自己的行为负责”
func main() {

	var age int
	fmt.Println("请输入年龄")

	// Scanln()函数通过age的地址(&age)
	// 找到变量age所在的空间，修改变量age存储空间里的值
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("你的年龄大于18，请为自己的行为负责")
	} else {
		fmt.Println("你的年龄不大，这次放过你了")
	}

	// Golang的特殊用法，支持在if中加入变量声明
	// if theAge := 30; theAge >= 18 {
	// 	fmt.Println("suck my dick jin yang")
	// }

	// if var theAge2 int = 30; theAge >= 18 {
	// 	fmt.Println("suck my dick jin yang")
	// }
}
