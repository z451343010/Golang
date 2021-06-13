/*
	牛客网
	华为机试
	HJ22
*/
package main

import "fmt"

func main() {

	var numSlice []int
	for {

		var bottleNum int
		fmt.Scanln(&bottleNum)
		if bottleNum == 0 {
			break
		}
		numSlice = append(numSlice, bottleNum)

	}

	for i := 0; i < len(numSlice); i++ {
		var count int = 0
		for {

			divide := numSlice[i] / 3
			count += divide
			remain := numSlice[i] % 3
			numSlice[i] = remain + divide
			if numSlice[i] < 3 {
				if numSlice[i] == 2 {
					count += 1
				}
				break
			}

		}

		fmt.Println(count)
	}
}
