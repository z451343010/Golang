package main

import "fmt"

// 判断一个年份是否是闰年
func main() {

	var year int
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)

	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				fmt.Println("moon year")
			} else {
				fmt.Println("no moon year")
			}
		} else {
			fmt.Println("moon year")
		}
	} else {
		fmt.Println("no moon year")
	}

}
