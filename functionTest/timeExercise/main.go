package main

import (
	"fmt"
	"strconv"
	"time"
)

// 统计代码的执行时间
func test3() {

	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}

}

func main() {

	// 在执行test3之前，先获取当前的unix时间戳
	start := time.Now().Unix()
	test3()
	end := time.Now().Unix()
	fmt.Println("执行test3()耗费的时间为：", (end - start), "秒")

}
