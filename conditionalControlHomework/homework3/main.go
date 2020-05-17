package main

import (
	"fmt"
	"strconv"
)

// 输入一个十进制数转二进制数
func main() {

	var tenNum int64
	var binStr string

	fmt.Println("请输入一个十进制数：")
	fmt.Scanln(&tenNum)

	for {

		var temp int64
		var remiander int64

		// 辗转相除，求余数，然后用字符串去拼接余数
		remiander = tenNum % 2 // 每次的余数
		temp = tenNum / 2
		tenNum = temp
		binStr += strconv.FormatInt(remiander, 10)

		if tenNum == 0 {
			break
		}

	}

	fmt.Println("转换后的二进制数是：")

	// 倒序输出原来的字符串
	for i := len(binStr) - 1; i >= 0; i-- {
		fmt.Printf("%c", binStr[i])
	}

}
