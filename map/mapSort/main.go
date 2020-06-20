package main

import (
	"fmt"
	"sort"
)

// map的排序
func main() {

	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	// 如果按照map的key的顺序进行排序输出
	// 1.先将map的key放入到切片中
	// 2.对切片排序
	// 3.遍历切片，然后按照key来输出map的值
	fmt.Println(map1)

	var key []int
	for k, _ := range map1 {
		key = append(key, k)
	}

	sort.Ints(key)
	fmt.Println(key)

	for _, v := range key {
		fmt.Println("map1[", v, "] = ", map1[v])
	}

}
