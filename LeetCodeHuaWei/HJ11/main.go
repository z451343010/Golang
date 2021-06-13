/*
	牛客网
	华为机试
	HJ11
*/
package main

import "fmt"

func main() {

	var num int
	fmt.Scanln(&num)

	for {

		remain := num % 10
		num = num / 10
		fmt.Print(remain)

		if num == 0 {
			break
		}

	}

}
