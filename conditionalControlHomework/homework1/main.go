package main

import "fmt"

// 实现判断一个整数，属于哪个范围
// 大于0；小于0；等于0
func main() {

	var intNum int64
	fmt.Println("请输入该整数：")
	fmt.Scanln(&intNum)

	if intNum > 0 {
		fmt.Println("大于0")
	} else if intNum < 0 {
		fmt.Println("小于0")
	} else if intNum == 0 {
		fmt.Println("等于0")
	} else {
		fmt.Println("输入有误")
	}

}
