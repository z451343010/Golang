package main

import "fmt"

// 演示管道 channel 的使用
func main() {

	// 创建一个可以存放3个int类型的channel
	var intChan chan int
	intChan = make(chan int, 3)

	// 输出 intChan
	fmt.Println(intChan)

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 看看管道的长度和cap（容量）
	fmt.Println("channel intChan len = ", len(intChan))
	fmt.Println("channel intChan cap = ", cap(intChan))

	// 从管道中取出数据
	// 管道每取出一个数据，管道长度-1
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)
	fmt.Println("channel intChan len = ", len(intChan))
	fmt.Println("channel intChan cap = ", cap(intChan))

}
