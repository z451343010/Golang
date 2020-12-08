/*
	数据结构
	稀疏数组
	1）使用稀疏数组，来保存类似前面的二维数组（棋盘、地图等等）
	2）把稀疏数组存盘分析，并且可以重新恢复原来的二维数组
*/
package main

import "fmt"

type Node struct {
	row    int // 行
	column int // 列
	value  int // 值
}

func main() {

	var chessBoard [11][11]int
	chessBoard[1][2] = 1 // 黑棋
	chessBoard[2][3] = 1 // 黑棋
	chessBoard[3][2] = 1 // 黑棋
	chessBoard[2][2] = 2 // 白棋

	// 输出原始的数组
	for _, v := range chessBoard {

		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}

		fmt.Println() // 换行

	}

	// 转成稀疏数组

	// 以切片的方式保留稀疏数组
	var sparseArray []Node

	// 初始化节点
	valNode := Node{
		row:    11,
		column: 11,
		value:  0,
	}

	sparseArray = append(sparseArray, valNode)

	// 转化数组为稀疏数组
	for i, v := range chessBoard {

		for j, v2 := range v {

			if v2 != 0 {

				valNode := Node{
					row:    i,
					column: j,
					value:  v2,
				}

				sparseArray = append(sparseArray, valNode)

			}

		}

	}

	// 输出稀疏数组
	fmt.Println(sparseArray)

}
