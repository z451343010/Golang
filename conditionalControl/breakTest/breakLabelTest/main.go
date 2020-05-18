package main

import "fmt"

// break label
// break语句出现在多层嵌套的语句块时
// 可以通过标签指明要终止的是哪一层语句块
func main() {

label2:
	for i := 0; i < 4; i++ {

		// lable1:
		for j := 0; j < 10; j++ {

			if j == 2 {
				// break
				// break lable1
				break label2
			}

			fmt.Println("j = ", j)

		}

	}

}
