package main

import "fmt"

/*
	编写方法，使得给定的一个二维数组（3*3）转置
*/
type Reverse struct {
}

func (reverse *Reverse) reverse(multiArr *[3][3]int) {

	var tempArr [3][3]int

	for i := 0; i < len(*multiArr); i++ {

		for j := 0; j < len((*multiArr)[i]); j++ {

			tempArr[i][j] = (*multiArr)[j][i]

		}

	}

	for i := 0; i < len(*multiArr); i++ {

		for j := 0; j < len((*multiArr)[i]); j++ {

			(*multiArr)[i][j] = tempArr[i][j]

		}

	}

	for i := 0; i < len(*multiArr); i++ {

		for j := 0; j < len((*multiArr)[i]); j++ {

			fmt.Print((*multiArr)[i][j], " ")

		}

		fmt.Println()

	}

}

func main() {

	var multiArr [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	var reverse Reverse
	(&reverse).reverse(&multiArr)

	fmt.Println("the multiArr of main()")
	fmt.Println(multiArr)

}
