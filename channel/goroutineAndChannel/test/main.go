package main

import "fmt"

// 测试管道中是否有默认值
func main() {

	exitChan := make(chan bool, 1)
	// exitChan <- true
	// fmt.Printf("%v\n", exitChan)
	_, ok := <-exitChan
	// fmt.Println("v = ", v)
	fmt.Println("ok = ", ok)

}
