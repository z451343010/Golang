/*
	将稀疏数组恢复成原始的数组
*/
package main

import "fmt"

type Node struct {
	row    int // 行
	column int // 列
	value  int // 值

}

func main() {

	// 稀疏数组
	var sparseArray []Node

	node0 := Node{
		row:    11,
		column: 11,
		value:  0,
	}

	node1 := Node{
		row:    1,
		column: 2,
		value:  1,
	}

	node2 := Node{
		row:    2,
		column: 2,
		value:  2,
	}

	node3 := Node{
		row:    2,
		column: 3,
		value:  1,
	}

	node4 := Node{
		row:    3,
		column: 2,
		value:  1,
	}

	sparseArray = append(sparseArray, node0)
	sparseArray = append(sparseArray, node1)
	sparseArray = append(sparseArray, node2)
	sparseArray = append(sparseArray, node3)
	sparseArray = append(sparseArray, node4)

	// 稀疏数组转化为二维数组
	var arr [11][11]int

	for i := 1; i < len(sparseArray); i++ {

		if sparseArray[i].value != 0 {

			arr[sparseArray[i].row][sparseArray[i].column] = sparseArray[i].value

		}

	}

	// 遍历二维数组
	for _, v := range arr {

		for _, v2 := range v {
			fmt.Print(v2, " ")
		}

		fmt.Println()

	}

}
