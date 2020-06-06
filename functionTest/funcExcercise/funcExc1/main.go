package main

import (
	"errors"
	"fmt"
)

// 函数综合课堂练习
// 循环打印输入的月份的天数。【使用continue实现】
// 要有判断输入的月份是否错误的语句

func isMoonYear(year int) bool {

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	} else {
		return false
	}

}

func judgeMonthDay(year int, month int, day int) error {

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:

		if day < 1 || day > 31 {
			fmt.Println("输入的日期有误")
			return errors.New("输入的日期有误")
		}

	case 4, 6, 9, 11:

		if day < 1 || day > 30 {
			fmt.Println("输入的日期有误")
			return errors.New("输入的日期有误")
		}

	default:

		if isMoonYear(year) {

			if day < 1 || day > 29 {
				fmt.Println("输入的日期有误")
				return errors.New("输入的日期有误")
			}

		} else {

			if day < 1 || day > 28 {
				fmt.Println("输入的日期有误")
				return errors.New("输入的日期有误")
			}

		}

	}

	fmt.Println("您输入的日期是:", year, "年",
		month, "月", day, "日")

	return nil

}

func main() {

	for {
		fmt.Println("请输入年:")
		var year int
		fmt.Scanln(&year)

		fmt.Println("请输入月:")
		var month int
		fmt.Scanln(&month)

		fmt.Println("请输入日:")
		var day int
		fmt.Scanln(&day)

		err := judgeMonthDay(year, month, day)

		if err != nil {
			continue
		}

	}

}
