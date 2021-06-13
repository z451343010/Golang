/*
	牛客网
	华为机试
	HJ15
*/
package main

import "fmt"

func main() {

	var num int
	fmt.Scanln(&num)

	if num == 0 {
		fmt.Println(0)
		return
	}

	var count int
	for num != 0 {

		remain := num % 2
		num = num / 2

		if remain == 1 {
			count++
		}

	}

	fmt.Print(count)

}
