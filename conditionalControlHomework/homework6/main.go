package main

import "fmt"

// 根据输入的年份和月份，判断当月的天数
func main() {

	var isMoonYear bool
	var year int
	var month int

	fmt.Println("请输入年份：")
	fmt.Scanln(&year)
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)

	// 判断年份是否为闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		isMoonYear = true
	}

	switch month {

	case 1, 3, 5, 7, 8, 10, 12:

		fmt.Println("该月份的天数为31天")

	case 2:

		if isMoonYear {
			fmt.Println("该月份的天数为29天")
		} else {
			fmt.Println("该月份的天数为28天")
		}

	default:

		fmt.Println("该月份的天数为30天")

	}

}
