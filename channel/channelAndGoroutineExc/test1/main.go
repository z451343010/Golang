package main

import "fmt"

/*
	利用管道(channel)和协程(goroutine)
	计算 1 到 n 的阶乘
*/

// 计算每个数的阶乘
func factorial(num int, mapChan chan map[int]int, countChan chan int, count int) {

	var result int = 1

	mapNum := make(map[int]int, 10)

	for i := 1; i <= num; i++ {
		result *= i
	}

	mapNum[num] = result
	fmt.Println("write working")
	mapChan <- mapNum

	// 计算写入次数的管道
	countChan <- num

	// 当写入的次数符合条件，即关闭管道
	if len(countChan) == count {
		close(countChan)
		close(mapChan)
	}

}

// 输出阶乘运算的结果
func readFactorial(mapChan chan map[int]int, exitChan chan bool) {

	for {
		v, ok := <-mapChan
		if !ok {
			break
		}
		fmt.Println("read working, v = ", v)
	}

	// 读完后，往exitChan写入一个标志
	// 说明该函数已经读完
	exitChan <- true
	close(exitChan)

}

func main() {

	// 当往管道中写入超过其容量（cap）的数据时，会报错
	// fatal error:all goroutines are asleep - deadlock!
	mapChan := make(chan map[int]int, 20000)
	exitChan := make(chan bool, 1)
	countChan := make(chan int, 20000)

	var count int = 1000 // 计算 1 到 count 的阶乘

	for i := 1; i <= count; i++ {
		// 开启 count 个协程，计算对应的阶乘
		go factorial(i, mapChan, countChan, count)
	}

	go readFactorial(mapChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
