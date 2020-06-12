package main

import (
	"fmt"
	"time"
)

// 二分查找（binary search）
func binarySearch(leftIndex int, rightIndex int,
	arr [6]int, searchNum int) {

	if leftIndex > rightIndex {
		fmt.Println("二分查找，没有找到")
		return
	}

	var middleIndex int
	middleIndex = (leftIndex + rightIndex) / 2

	if arr[middleIndex] == searchNum {

		fmt.Println("二分查找，已经找到，下标为：", middleIndex)
		return

	} else if arr[middleIndex] > searchNum {

		binarySearch(leftIndex, middleIndex-1, arr, searchNum)

	} else {

		binarySearch(middleIndex+1, rightIndex, arr, searchNum)

	}

}

// 顺序查找
func orderSearch(arr [6]int, searchNum int) {

	var flag bool

	for i := 0; i < len(arr); i++ {

		if arr[i] == searchNum {

			fmt.Println("顺序查找已经找到，下标为：", i)
			flag = true
			return

		}

	}

	if !flag {
		fmt.Println("顺序查找，没有找到")
	}

}

func main() {

	arr := [6]int{1, 8, 10, 89, 1000, 1234}

	var searchNum int
	fmt.Println("请输入要查找的数字:")
	fmt.Scanln(&searchNum)

	fmt.Println()

	binaryStartTime := time.Now().UnixNano()
	binarySearch(0, len(arr)-1, arr, searchNum)
	binaryEndTime := time.Now().UnixNano()

	// fmt.Println("--------")
	// fmt.Println("binaryStartTime", binaryStartTime)
	// fmt.Println("binaryEndTime", binaryEndTime)
	// fmt.Println("time sub", (binaryEndTime - binaryStartTime))
	// fmt.Println("--------")

	fmt.Println("二分查找的时间为:",
		(binaryEndTime - binaryStartTime), "纳秒")

	fmt.Println()

	orderStartTime := time.Now().UnixNano()
	orderSearch(arr, searchNum)
	orderEndTime := time.Now().UnixNano()

	fmt.Println("顺序查找的时间为:",
		(orderEndTime - orderStartTime), "纳秒")

}
