package main

import "fmt"

/*
	编写结构体（MethodUtils），编写一个方法
	方法不需要参数，在方法中打印一个 10*8 的矩形
	在 main() 中调用该方法
*/

type MethodUtils struct {
}

func (methodUtils MethodUtils) printRectangle() {

	for i := 0; i < 10; i++ {

		for j := 0; j < 8; j++ {

			fmt.Print("*")

		}

		fmt.Println()

	}

}

func main() {

	var methodUtils MethodUtils
	methodUtils.printRectangle()

}
