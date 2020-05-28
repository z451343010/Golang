package main

import "fmt"

// 在Go中支持可变参数
// 编写一个函数sum，可以求出1到多个int的和
func sum(n1 int, args ...int) int {

	var sum int = n1

	for i := 0; i < len(args); i++ {

		sum += args[i]

	}

	return sum

}

func main() {

	var sum int = sum(10, 20, 30, 40)
	fmt.Println("sum =", sum)

}
