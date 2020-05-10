package main

import "fmt"

// 课堂练习1
// 还有97天放假，问：xx个星期零xx天
func main() {
	var day int = 97
	var weekDays int = 7
	var remainDay int = day % weekDays
	var weekNum = day / weekDays
	fmt.Println(weekNum, "个星期零", remainDay, "天")
}
