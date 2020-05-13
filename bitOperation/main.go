package main

import "fmt"

func main() {

	// 位运算
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)
	fmt.Println(-2 ^ 2)

	// 移位运算
	var leftMove int = 1 << 2  // 0000 0001 => 0000 0100 = 4
	var rightMove int = 1 >> 2 // 0000 0001 => 0000 0000 = 0

	fmt.Println(leftMove)
	fmt.Println(rightMove)

}
