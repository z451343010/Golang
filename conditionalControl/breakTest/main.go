package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成一个1-100的一个数
// 直到生成了99这个数
// 数数看一共用了几次
func main() {

	var count int
	var randNum int

	for {

		// 我们为了生成一个随机数，还需要设置一个种子(类似于随机源)
		// time.Now().Unix()
		// Unix时间，从1970年1月1日0时0分0秒起至现在的总微秒数
		rand.Seed(time.Now().UnixNano())

		// rand.Intn() 生成[0,100)的随机数
		// 在Golang中，如果没有随机种子 rand.Seed()
		// 每次生成的随机数将会是相同的
		randNum = rand.Intn(100 + 1)

		count++
		if randNum == 99 {
			break
		}

	}

	fmt.Println("一共生成了：", count, "次")

}
