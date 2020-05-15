package main

import "fmt"

// 课堂练习
// 出票系统：根据淡旺季的月份和年龄，打印票价
// 4_10 旺季
// 	成人（18-60）：60
//  儿童（<18）：
//  老人 （>60）：1/2
// 淡季
//  成人：40
//  其他：20
func main() {

	var month int8
	var age int8
	var price float64

	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 { // 旺季
		if age >= 18 && age <= 60 {
			price = 60
		} else if age < 18 {
			price = 60 / 2
		} else {
			price = 60 / 3
		}
	} else {
		if age >= 18 && age <= 60 {
			price = 40
		} else {
			price = 20
		}
	}

	fmt.Println("票价为：", price)

}
