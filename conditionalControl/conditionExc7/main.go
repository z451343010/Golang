package main

import (
	"fmt"
	"math"
)

// 求解一元二次方程 a*x^2 + bx + c = 0
func main() {

	var a float64 = 1
	var b float64 = -4
	var c float64 = 4
	var x1 float64
	var x2 float64

	if b*b-4*a*c >= 0 {

		x1 = (-b + math.Sqrt(b*b-4*a*c)) / 2 * a
		x2 = (-b - math.Sqrt(b*b-4*a*c)) / 2 * a

		if x1 == x2 {
			fmt.Println("该方程的解是：", x1)
		} else {
			fmt.Println("该方程有两个解，一个是：", x1, "另一个为：", x2)
		}

	} else {
		fmt.Println("该方程无解")
	}

}
