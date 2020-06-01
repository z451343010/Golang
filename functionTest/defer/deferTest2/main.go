package main

import "fmt"

// 在 defer 将语句放入到栈的时候
// 也会将相关的值拷贝同时入栈
func sum(n1 int, n2 int) {

	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n2 = ", n2)

	var result int = n1 + n2
	n1++
	fmt.Println("ok3 result = ", result, "n1 = ", n1)

}

func main() {

	sum(10, 20)

}
