/*
	复习
	100 以内的数求和，求出当和第一次大于 20 的当前数
*/
package main

import "fmt"

func main() {

label1:
	for i := 1; i <= 100; i++ {

		for j := 1; j <= 100; j++ {

			if (i + j) > 20 {
				fmt.Println("i = ", i)
				fmt.Println("j = ", j)
				break label1
			}

		}

	}

}
