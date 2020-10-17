/*
	常量专业的写法
*/
package main

import "fmt"

func main() {

	/*
		表示给 a 赋值为 0
		b 在 a 的基础上+1
		c 在 a 的基础上+1
	*/
	const (
		a = iota
		b
		c
	)

	fmt.Println(a, b, c)

}
