package main

import "fmt"

/*
	管道可以声明为只读或者只写
*/
func main() {

	// 在默认情况下，管道是双向的
	// var chan1 chan int  // 可读可写

	// 声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 10)
	chan2 <- 20
	fmt.Println("chan2 = ", chan2)

	var chan3 <-chan int
	chan3 = make(chan int, 10)
	num2 := <-chan3
	fmt.Println("num2", num2)

}
