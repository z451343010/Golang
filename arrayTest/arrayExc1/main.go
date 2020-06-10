package main

import "fmt"

// 创建一个byte类型的26个元素的数组
// 分别放置'A'-'Z'
// 使用for循环遍历所有元素并打印出来
func main() {

	var charArr [26]byte

	for i := 0; i < 26; i++ {

		charArr[i] = byte('A' + i)
		fmt.Printf("%c ", charArr[i])

	}

}
