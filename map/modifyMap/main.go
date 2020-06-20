package main

import "fmt"

// map是引用类型，遵守引用类型传递的机制
// 在一个函数接收map，修改后，会直接修改原来的值

func modify(map1 map[int]int) {
	map1[1] = 900
}

func main() {

	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 88
	map1[3] = 1
	map1[4] = 2
	modify(map1)
	fmt.Println(map1)

}
