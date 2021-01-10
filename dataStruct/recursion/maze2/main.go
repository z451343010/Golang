/*
	数据结构
	递归
	迷宫问题 求最短路径
*/
package main

import "fmt"

type Box struct {
	x int
	y int
}

type Path struct {
	data   [200]Box
	length int
}

func mazePath(startX int, startY int, endX int, endY int,
	path *Path, mazeMap *[10][10]int) {

	var (
		direction int
		x         int
		y         int
	)

	if startX == endX && startY == endY {
		path.data[path.length].x = startX
		path.data[path.length].y = startY
		path.length++

		for i := 0; i < path.length; i++ {
			fmt.Printf("(%d,%d)\n", path.data[i].x, path.data[i].y)
		}

	} else {

		if mazeMap[startX][startY] == 0 {

			direction = 0
			for direction < 4 {

				path.data[path.length].x = startX
				path.data[path.length].y = startY
				path.length++

				switch direction {
				// 上
				case 0:
					x = startX - 1
					y = startY
				// 右
				case 1:
					x = startX
					y = startY + 1
				// 下
				case 2:
					x = startX + 1
					y = startY
				// 左
				case 3:
					x = startX
					y = startY - 1
				}

				mazeMap[startX][startY] = -1
				mazePath(x, y, endX, endY, path, mazeMap)
				mazeMap[startX][startY] = 0

				path.length--
				direction++

			}

		}

	}

}

func main() {

	mazeMap := [10][10]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 1, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 1, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 0, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	// 输出地图
	fmt.Println("==========迷宫地图==========\n")
	for i := 0; i < len(mazeMap); i++ {

		for j := 0; j < len(mazeMap[i]); j++ {
			fmt.Print(mazeMap[i][j], " ")
		}

		fmt.Println()

	}

	// 走迷宫
	var (
		startX int = 1
		startY int = 1
		endX   int = 8
		endY   int = 8
	)

	path := &Path{}

	mazePath(startX, startY, endX, endY, path, &mazeMap)

	fmt.Println(path)
}
