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
	data   [50]Box
	length int
}

func mazePath(startX int, startY int, endX int, endY int,
	path *Path, mazeMap *[4][4]int) {

	var (
		direction int
		x         int
		y         int
	)

	if startX == endX && startY == endY {
		path.data[path.length].x = startX
		path.data[path.length].y = startY
		path.length++
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

	var mazeMap [4][4]int

	// 把最外的一圈设置为墙
	for i := 0; i < len(mazeMap); i++ {

		for j := 0; j < len(mazeMap[i]); j++ {
			if i == 0 || i == (len(mazeMap)-1) {
				mazeMap[i][j] = 1
			}
			if j == 0 || j == (len(mazeMap[i])-1) {
				mazeMap[i][j] = 1
			}
		}

	}

	// mazeMap[3][1] = 1
	// mazeMap[3][2] = 1

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
		endX   int = 2
		endY   int = 2
	)

	path := &Path{}

	mazePath(startX, startY, endX, endY, path, &mazeMap)

	fmt.Println(path.data[0])

}
