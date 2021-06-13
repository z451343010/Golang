/*
	牛客网
	华为机试
	HJ55
*/
package main

import "fmt"

func main() {

	for j := 0; j < 2; j++ {
		var num int
		fmt.Scanf("%d\n", &num)

		count := 0
		for i := 6; i <= num; i++ {

			flag := false

			if i%7 == 0 {
				count++
			}

			tempNum := i
			for tempNum != 0 {
				remain := tempNum % 10
				divide := tempNum / 10
				tempNum = divide

				if remain == 7 {
					flag = true
					break
				}
			}

			if flag && i%7 != 0 {
				count++
			}

		}
		fmt.Println(count)
	}

}
