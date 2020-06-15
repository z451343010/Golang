package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组和排序的作业1
// 随机生成10个整数[1,100]保存到数组
// 并倒序打印
// 以及求平均值、最大值和最大值的下标
// 查找里面是否有55

// 冒泡排序（倒序）
func bubbleSort(arr *[10]int) {

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-1-i; j++ {

			if arr[j] < arr[j+1] {

				var temp int
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp

			}

		}

	}

	fmt.Println("倒序打印后的数组为：")
	for k := 0; k < len(arr); k++ {
		fmt.Print(arr[k], " ")
	}

}

// 求平均值、最大值、最大值的下标
func aveAndMaxAndMaxIndex(arr [10]int) (ave float64, max int, maxIndex int) {

	var sum int
	max = arr[0]

	for i := 1; i < len(arr); i++ {

		sum += arr[i]

		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}

	ave = float64(sum) / float64(len(arr))

	return ave, max, maxIndex

}

func main() {

	rand.Seed(time.Now().UnixNano())

	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100) + 1
	}

	fmt.Println(arr)
	ave, max, maxIndex := aveAndMaxAndMaxIndex(arr)
	fmt.Println("ave = ", ave)
	fmt.Println("max = ", max)
	fmt.Println("maxIndex = ", maxIndex)
	bubbleSort(&arr)

}
