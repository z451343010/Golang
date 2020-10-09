/*
	作业1：goroutine课堂练习
	要求：
	1)启动一个协程，将1-2000的数放入到一个channel中，比如numChan
	2)启动8个协程，从numChan取出数字（比如n），
	并计算1+..+n的值，并存放到resChan
	3)最后8个协程协同完成工作后，再遍历resChan，
	显示结果如[res[1]=1...res[10]=55...]
	4)注意：考虑resChan chan int 是否合适
*/
package main

import "fmt"

// 存放数据到 numChan
func putNum(numChan chan int, countNum int) {

	for i := 1; i <= countNum; i++ {
		fmt.Printf("put %d into numChan\n", i)
		numChan <- i
	}

	// 关闭管道 numChan
	close(numChan)

}

// 从numChan取出数字，统计累加的值
func accumulation(numChan chan int, resChan chan map[int]int, exitChan chan bool, countNum int) {

	for {

		v, ok := <-numChan
		var result int = 0

		if !ok {
			break
		}

		for i := 1; i <= v; i++ {
			result += i
		}

		// 把累加后的结果放入到 numMap
		numMap := make(map[int]int, 10)
		numMap[v] = result
		// numMap 放入到管道 resChan
		resChan <- numMap

		fmt.Printf("res[%v]=%v\n", v, result)

	}

	// 每个结果都计算完毕，关闭 reshChan
	if len(resChan) == countNum {
		close(resChan)
		// 把结束计算的标志放置到 exitChan
		exitChan <- true
		close(exitChan)
	}

}

func main() {

	numChan := make(chan int, 3000)
	resChan := make(chan map[int]int, 3000)
	exitChan := make(chan bool, 1)

	// 放入 1 到 countNum
	var countNum int = 2000

	// 开启存放整数的协程 putNum
	go putNum(numChan, countNum)

	for i := 0; i < 8; i++ {
		// 开启8个计算累加的协程 accumulation
		go accumulation(numChan, resChan, exitChan, countNum)
	}

	// 判断所有的协程是否执行完毕，即累加的计算全部完成
	for {

		_, ok := <-exitChan

		if !ok {
			break
		}

	}

}
