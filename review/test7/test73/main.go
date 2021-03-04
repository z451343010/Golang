/*
	复习
	switch 语句还可以用于 type-switch 来判断某个 interface 变量中
	实际指向的变量类型
*/
package main

import "fmt"

func main() {

	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Println("x 的类型：%T", i)
	case int:
		fmt.Println("x 是 int 类型")
	case float64:
		fmt.Println("x 是 float64 类型")
	case float64:
		fmt.Println("x 是 float64 类型")
	}

}
