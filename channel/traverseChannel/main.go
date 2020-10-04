package main

import "fmt"

/*
	管道的遍历
*/
func main() {

	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	// 遍历前需要关闭管道，否则会出现死锁
	close(intChan2)

	// for range 遍历管道
	// 管道没有下标
	for value := range intChan2 {
		fmt.Println("value = ", value)
	}

}
