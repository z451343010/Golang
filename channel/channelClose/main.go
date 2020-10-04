package main

import "fmt"

func main() {

	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	// close()关闭管道
	// 这时不能再写入数据到该管道
	close(intChan)

	fmt.Println("ok")

	// 当管道关闭后，依旧可以读取数据
	n1 := <-intChan
	fmt.Println(n1)

}
