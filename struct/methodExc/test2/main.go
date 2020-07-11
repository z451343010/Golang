package main

import "fmt"

/*
	编写一个方法，提供m和n两个参数
	方法中打印一个 m * n 的矩形
*/

type MethodUtils struct {
}

func (methodUtils MethodUtils) printRectangle(m int, n int) {

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {

			fmt.Print("*")

		}

		fmt.Println()

	}

}

func main() {

	var methodUtils MethodUtils
	var m int
	var n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)

	methodUtils.printRectangle(m, n)

}
