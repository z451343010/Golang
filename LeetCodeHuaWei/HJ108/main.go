/*
	牛客网
	华为机试
	HJ108
	求两个数的最小公倍数
*/
package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Scanf("%d %d\n", &num1, &num2)
	var max int
	var min int
	if num1 >= num2 {
		max = num1
		min = num2
	} else {
		max = num2
		min = num1
	}

	if max%min == 0 {
		fmt.Print(max)
	} else {
		for i := max; i <= max*min; i++ {
			if i%max == 0 && i%min == 0 {
				fmt.Print(i)
				break
			}
		}
	}

}
