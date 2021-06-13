/*
	牛客网
	华为机试
	HJ6
	程序超时
*/
package main

import "fmt"

func main() {

	var num int
	fmt.Scanln(&num)

	for i := 2; i*i <= num; i++ {

		for num%i == 0 {
			fmt.Print(i, " ")
			num /= i
		}

	}

}
