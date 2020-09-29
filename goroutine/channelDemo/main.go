package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	计算1-200的各个数的阶乘，并把各个数的阶乘放入到map中。
	最后显示出来。要求使用goroutine完成。
*/

var (
	myMap = make(map[int64]int64, 10)
	// 声明一个全局的互斥锁
	// lock是一个全局的互斥锁
	// sync:synchornized 同步地
	// Mutex 互斥
	lock sync.Mutex
)

// 计算 n 的阶乘
func test(n int64) {

	res := int64(1)
	for i := int64(1); i <= n; i++ {
		res *= i
	}

	// 执行共有的资源之前，先加锁
	// Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁
	lock.Lock()
	myMap[n] = res
	// 访问完毕，解锁
	lock.Unlock()

}

func main() {

	for i := int64(1); i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()
	for index, value := range myMap {
		fmt.Printf("myMap[%d] = %d\n ", index, value)
	}
	lock.Unlock()

}
