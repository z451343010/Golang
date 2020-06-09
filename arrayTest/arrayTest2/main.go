package main

import "fmt"

func main() {

	var intArr [3]int
	// 当我们定义完数组后，数组的各个元素有默认值0
	fmt.Println("intArr[0]的地址：", &intArr[0])
	fmt.Println("intArr[1]的地址：", &intArr[1])

}
