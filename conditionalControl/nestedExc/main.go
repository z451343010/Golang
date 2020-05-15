package main

import "fmt"

// 嵌套分支课堂练习
// 判断某人是否进入决赛，以及进入男子组还是女子组
func main() {

	var second float64
	var gender string

	fmt.Println("请输入时间（秒）：")
	fmt.Scanln(&second)

	fmt.Println("请输入性别：")
	fmt.Scanln(&gender)

	if second <= 8 {

		if gender == "男" {
			fmt.Println("进入男子组")
		} else if gender == "女" {
			fmt.Println("进入女子组")
		} else {
			fmt.Println("输入的性别无效")
		}

	} else {
		fmt.Println("淘汰")
	}

}
