package main

import "fmt"

// 用continue实现打印1到100的所有奇数
func main() {

	for i := 1; i <= 100; i++ {

		if i%2 == 0 {
			continue
		}

		fmt.Println(i)

	}

}
