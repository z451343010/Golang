package main

import "fmt"

// 根据用户指定的月份，打印出属于哪个季节
func main() {

	var month int8

	fmt.Println("请输入月份：")
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println(month, "月，是春季")
	case 6, 7, 8:
		fmt.Println(month, "月，是夏季")
	case 9, 10, 11:
		fmt.Println(month, "月，是秋季")
	case 12, 1, 2:
		fmt.Println(month, "月，是冬季")
	default:
		fmt.Println("输入的月份无效")
	}

}
