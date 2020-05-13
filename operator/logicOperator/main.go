package main

import "fmt"

// 演示逻辑运算符
func main() {

	// 逻辑与&& 的使用
	var age int = 40

	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	// 逻辑或|| 的使用
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}

	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	// 逻辑非！的使用
	if !(age > 30) {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}

}
