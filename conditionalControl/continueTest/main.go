package main

import "fmt"

// continue example
func main() {

label1:
	for i := 0; i < 4; i++ {
		// label2:
		for j := 0; j < 10; j++ {

			if j == 2 {
				continue label1
			}

			fmt.Println("j = ", j)

		}

		fmt.Println()

	}

}
