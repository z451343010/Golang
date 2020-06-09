package main

import "fmt"

// 循环输入5个成绩，保存到float64数组，并输出
func main() {

	var floatArr [5]float64

	for i := 0; i < len(floatArr); i++ {
		fmt.Println("输入数组中的第", i, "个元素：")
		fmt.Scanln(&floatArr[i])
	}

	fmt.Println("循环输出数组：")

	for j := 0; j < len(floatArr); j++ {
		fmt.Println(floatArr[j])
	}

}
