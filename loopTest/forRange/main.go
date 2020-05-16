package main

import "fmt"

// 使用 for range 方式遍历字符串
// index, val 类似于键值对
// 可以把 index, val 看成是一个整体
func main() {

	var testStr string = "Hello,World!北京"

	for index, val := range testStr {
		fmt.Printf("index = %d,val = %c\n", index, val)
	}

}
