package main

import "fmt"

func main() {
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // 穿透，无论下一个语句的条件是否成立
		// 直接执行下一个语句里面的代码块
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("ok4")
	}
}
