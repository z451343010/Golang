package main

import (
	"fmt"
	"runtime"
)

func main() {

	// 获取当前的cpu数量
	num := runtime.NumCPU()
	fmt.Println(num)

	// 这里设置 num - 1 的cpu运行go程序
	runtime.GOMAXPROCS(num)

	fmt.Println("num = ", num)

}
