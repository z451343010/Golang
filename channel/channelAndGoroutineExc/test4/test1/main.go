package main

import "fmt"

func main() {

	var flag bool = true

	for i := 1; i < 10; i++ {

		fmt.Println(" i = ", i)
		fmt.Println("flag = ", flag)

		for j := 2; j < i; j++ {

			if i%j == 0 {
				flag = false
				break
			}

		}

		if flag {
			fmt.Println(i)
		}

	}

}
