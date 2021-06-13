/*
	牛客网
	华为机试
	HJ8
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	var putNum int
	fmt.Scanln(&putNum)

	var numMap map[int]int
	numMap = make(map[int]int, putNum)

	for i := 0; i < putNum; i++ {
		var index int
		var value int
		fmt.Scanf("%d %d\n", &index, &value)
		numMap[index] += value
	}

	// 先把 map 的 key 保存到一个切片
	// 对切片进行排序
	var indexs []int
	for index := range numMap {
		indexs = append(indexs, index)
	}
	sort.Sort(sort.IntSlice(indexs))

	// 遍历切片，根据切片里保存的 key 来反向输出 map 里的value
	for _, value := range indexs {
		fmt.Println(value, numMap[value])
	}

}
