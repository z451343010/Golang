package main

import "fmt"

func main() {
	var i int = 9
	// 二进制输出
	fmt.Printf("i = %b \n", i)

	// 八进制
	var j int = 012
	fmt.Println("j =", j)

	// 十六进制
	var k int = 0x18
	fmt.Println("k =", k)

}
