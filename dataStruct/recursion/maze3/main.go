/*
	数据结构
	顺序栈
	走迷宫算法
	《数据结构教程（第4版）》 P77
*/
package main

import "fmt"

type Box struct {
	x         int // 横坐标
	y         int // 纵坐标
	direction int // 方向
}

type Stack struct {
	data [50]Box
	top  int
}

func mazePath(startX int, startY int, endX int, endY int,
	path *Stack, mazeMap *[10][10]int) bool {

	var (
		x         int
		y         int
		direction int // 方向
		find      int // 标识是否可走
	)

	path.top = -1 // 初始化栈顶元素
	// 将起点压入栈
	path.top++
	path.data[path.top].x = startX
	path.data[path.top].y = startY
	// 初始化走相邻格子的方向
	path.data[path.top].direction = -1
	mazeMap[startX][startY] = -1

	for path.top > -1 {

		// 每次将格子坐标、方向压入栈
		// 不可走的格子坐标、方向后续将会被弹出栈
		x = path.data[path.top].x
		y = path.data[path.top].y
		direction = path.data[path.top].direction

		// 走到了终点，结束程序
		// 返回函数执行结果为 true
		if x == endX && y == endY {
			return true
		}

		// 某方向的相邻格子是否能走
		// 如果能走，find 置为 1
		find = 0

		for direction < 4 && find == 0 {

			direction++
			// 探索相邻的格子
			// direction = 0 向上探索
			// direction = 1 向右探索
			// direction = 2 向下探索
			// direction = 3 向左探索
			switch direction {
			case 0:
				x = path.data[path.top].x - 1
				y = path.data[path.top].y
			case 1:
				x = path.data[path.top].x
				y = path.data[path.top].y + 1
			case 2:
				x = path.data[path.top].x + 1
				y = path.data[path.top].y
			case 3:
				x = path.data[path.top].x
				y = path.data[path.top].y - 1
			}

			// 说明该方向的格子可以走
			// 把 find 置为 1，结束搜索方向的 for 循环
			if mazeMap[x][y] == 0 {
				find = 1
			}

		}

		// 如果该方向可以走，则把该方向的格子压入栈
		if find == 1 {
			path.data[path.top].direction = direction
			path.top++
			path.data[path.top].x = x
			path.data[path.top].y = y
			path.data[path.top].direction = -1
			mazeMap[x][y] = -1
		} else { // 如果不能走，则把该方向的格子弹出栈
			mazeMap[path.data[path.top].x][path.data[path.top].y] = 0
			path.top--
		}

	}

	// 如果走不到终点，说明该迷宫问题无解
	return false

}

func main() {

	// 初始化迷宫地图
	// 用二维数组来表示
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

	var (
		startX int = 1 // 起点横坐标
		startY int = 1 // 起点纵坐标
		endX   int = 8 // 终点横坐标
		endY   int = 8 // 终点纵坐标
	)

	// 初始化栈
	path := &Stack{}

	// 如果迷宫问题有解
	// 输出走迷宫的步骤
	if mazePath(startX, startY, endX, endY, path, &mazeMap) {

		for i := 0; i <= path.top; i++ {
			var step int = i + 1
			fmt.Printf("第 %d 步  ===>  (%d,%d)\n",
				step, path.data[i].x, path.data[i].y)
		}

	} else {
		fmt.Println("该迷宫问题无解")
	}

}
