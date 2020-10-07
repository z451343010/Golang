package main

import "fmt"

/*
	goroutine和channel结合应用实例
	请完成goroutine和channel协同工作的案例，具体要求：
	1）开启一个writeData协程，向管道intChan中写入50个整数
	2）开启一个readData协程，从管道intChan中读取writeData写入的数据。
	3）注意：writeData和readData操作的是同一个管道。
	4）主线程需要等待writeData和readData协程都完成工作才能退出
*/

// write Data
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println(" ====== writeData working ======")
	}
	// 数据写完进管道，即可关闭
	close(intChan)
}

// read Data
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(" ====== readData working ======")
		fmt.Println("readData read data = ", v)
	}

	// readData 读取完数据以后，即任务完成
	exitChan <- true
	close(exitChan)

}

func main() {

	// 创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
