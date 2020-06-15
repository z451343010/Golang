package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成10个整数(1-100)
// 使用冒泡排序法进行排序
// 使用二分查找法进行查找
// 查找是否有90这个数，并显示下标
// 没有则提示找不到这个数

// 二分查找
func binarySearch(arr [10]int, leftIndex int, rightIndex int,
	searchNum int) {

	if leftIndex > rightIndex {
		fmt.Println("找不到该数")
		return
	}

	var middleIndex int = (leftIndex + rightIndex) / 2

	if searchNum < arr[middleIndex] {
		binarySearch(arr, leftIndex, middleIndex-1, searchNum)
	} else if searchNum > arr[middleIndex] {
		binarySearch(arr, middleIndex+1, rightIndex, searchNum)
	} else {
		fmt.Println("该数的下标为 = ", middleIndex)
	}

}

// 冒泡排序
func bubbleSort(arr *[10]int) {

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				var temp int
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

}

func main() {

	// 随机生成10个整数
	var numArr [10]int
	// 根据Unix时间戳生成随机源
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(numArr); i++ {
		numArr[i] = rand.Intn(100) + 1
	}

	fmt.Println("排序前的数组为：", numArr)

	// 冒泡排序
	bubbleSort(&numArr)
	fmt.Println("排序后的数组为：", numArr)

	// 二分查找
	binarySearch(numArr, 0, len(numArr)-1, 90)

}
