package main

import "fmt"

// 输出三个整数进行排序
// 输出时按从小到大的顺序排序
func main() {

	var a int
	var b int
	var c int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	if a > b {
		if b > c {
			fmt.Println(c, " ", b, " ", a, " ")
		} else if b < c && a < c {
			fmt.Println(b, " ", a, " ", c, " ")
		} else {
			fmt.Println(b, " ", c, " ", a, " ")
		}
	} else {

		if c > b {
			fmt.Println(a, " ", b, " ", c, " ")
		} else if c > a && a < b {
			fmt.Println(a, " ", c, " ", b, " ")
		} else {
			fmt.Println(c, " ", a, " ", b, " ")
		}

	}

}
