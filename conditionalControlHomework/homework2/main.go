package main

import "fmt"

// 水仙花数：一个3位数，各个位上数字的立方和等于本身
// 如：153 = 1*1*1 + 5*5*5 + 3*3*3
// 输出所有的水仙花数
func main() {

	for i := 100; i < 1000; i++ {

		var sum int
		var temp int
		var remiander int

		var tempI int = i

		for {

			// 辗转相除，求各个位上的值
			remiander = tempI % 10
			temp = tempI / 10
			tempI = temp

			// 累加各个位上的立方
			sum += remiander * remiander * remiander

			if tempI == 0 {
				break
			}

		}

		// 判断各个位数的立方和是否等于本身
		// 如果相等，判断为水仙花数
		if sum == i {
			fmt.Println(i, "是水仙花数")
		}

	}

}
