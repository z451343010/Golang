/*
	goroutine和channel结合
	应用实例3
	需求：要求统计1-200000的数字中，哪些是素数？
	[为了缩短测试时间，把测试数据200000改为80000]
*/
package main

import "fmt"

// 向管道中放入数字
func putNum(intChan chan int, countNum int) {

	for i := 1; i <= countNum; i++ {
		fmt.Println("往管道中放入数据:", i)
		intChan <- i
	}

	close(intChan)

}

// 读取管道中的数字，并判断是否是素数
func readPrime(intChan chan int, primeChan chan int,
	exitChan chan bool) {

	var flag bool

	for {

		v, ok := <-intChan
		if !ok {
			break
		}

		flag = true

		for i := 2; i < v; i++ {

			if v%i == 0 {
				flag = false
				break
			}

		}

		if flag {
			primeChan <- v
		}

	}

	exitChan <- true

}

func main() {

	var routineNum int // 开启的协程数量
	var countNum int   // 统计 1-countNum的数字

	// 设置开启的协程数量、countNum
	routineNum = 12
	countNum = 80000

	intChan := make(chan int, countNum)
	primeChan := make(chan int, 80000)      // 放入质数的结果
	exitChan := make(chan bool, routineNum) // 标识退出的管道

	// 开启一个协程，向 intChan 放入 1-countNum 个数
	go putNum(intChan, countNum)

	// 开启 routineNum 个协程，判断是否为质数
	for i := 0; i < countNum; i++ {
		go readPrime(intChan, primeChan, exitChan)
	}

	// 开启协程，查看是否从 exitChan 里读取了 routineNum 个数据
	// 如果没有读取到 routineNum 个数据，则协程处于阻塞状态
	// 直到从 exitChan 读取了相应数量的数据，阻塞才解除
	go func() {

		for i := 0; i < routineNum; i++ {
			<-exitChan
		}

		close(primeChan)
		// close(exitChan)

	}()

	fmt.Printf("1 到 %d 素数为：\n", countNum)

	for {

		result, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(result)

	}

}
