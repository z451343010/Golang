package main

import "fmt"

/*
	在MethodUtils结构体编写个方法
	从键盘接收整数（1-9），打印对应乘法表
*/

type MethodUtils struct {
	num int
}

func (methodUtils *MethodUtils) printTable() {

	for i := 1; i <= methodUtils.num; i++ {

		for j := 1; j <= i; j++ {

			fmt.Print(j, "*", i, "=", j*i, "\t")

		}

		fmt.Println()

	}

}

func main() {

	var num int
	fmt.Scanln(&num)

	var methodUtils MethodUtils
	(&methodUtils).num = num
	(&methodUtils).printTable()

}
