/*
	牛客网
	华为机试
	HJ9
*/
package main

import "fmt"

func main() {

	var num int
	fmt.Scanln(&num)

	var numSlice []int

	for num%10 != 0 {

		remain := num % 10
		num = num / 10

		if len(numSlice) == 0 {
			numSlice = append(numSlice, remain)
		} else {

			// 去除重复的数字
			var flag bool

			for _, value := range numSlice {
				if remain == value {
					flag = true
					break
				}
			}

			if !flag {
				numSlice = append(numSlice, remain)
			}

		}

	}

	for _, value := range numSlice {
		fmt.Print(value)
	}

}
