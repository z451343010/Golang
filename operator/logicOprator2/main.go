package main

import "fmt"

func test() bool {

	fmt.Println("test...")
	return true
}

func main() {

	var i int = 10

	// 逻辑与，因为第一个判断条件为false
	// 所以运算结果为false，所以后面语句的不执行
	if i < 9 && test() {
		fmt.Println("ok...")
	}

	// 逻辑或，因为第一个判断条件为true
	// 所以运算的结果为true，后面语句的不执行
	if i > 9 || test() {
		fmt.Println("ok...")
	}

}
