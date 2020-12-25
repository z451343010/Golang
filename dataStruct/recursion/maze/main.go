/*
	数据结构
	递归
	走迷宫算法
	1）因为没有学习Go语言的GUI开发-GTK，因此我们在后台完成迷宫问题。
	2）小球得到的路径，和程序员设置的找路策略有关，
	即：找路的上下左右的顺序相关。
	3）再得到小球路径时，可以先使用（下右上左），
	再改成（上右下左），看看路径是不是有变化。
	4）思考：如何求出最短路径？
*/
package main

import "fmt"

// 找迷宫的出路的算法
// mazeMap 迷宫的地图
// i，j 出发点的坐标
func GetMazeWay(mazeMap *[8][7]int, i int, j int) bool {

	if mazeMap[6][5] == 2 {
		return true
	} else {

		// 说明要继续找
		// 说明可以继续探索，起始点不在墙上
		if mazeMap[i][j] == 0 {

			// 假设这个点是通路
			// 但是能不能通还需要具体探测
			// 探测的策略 （下、右、上、左）
			mazeMap[i][j] = 2
			if GetMazeWay(mazeMap, i+1, j) { // 向下走
				return true
			} else if GetMazeWay(mazeMap, i, j+1) { // 向右走
				return true
			} else if GetMazeWay(mazeMap, i-1, j) { // 向上走
				return true
			} else if GetMazeWay(mazeMap, i, j-1) { // 向左走
				return true
			} else { // 死路
				mazeMap[i][j] = 3
				return false
			}

		} else { // 说明这个点不能探测，是一堵墙
			return false
		}

	}

}

func main() {

	// 先创建一个二维数组，模拟一个迷宫地图
	// 1.代表墙
	// 0.代表没有走过的点
	// 2.元素的值为2，表示可以走通的路
	// 3.表示曾经走过，但是是一条死路
	var mazeMap [8][7]int

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

	mazeMap[3][1] = 1
	mazeMap[3][2] = 1

	// 输出地图
	fmt.Println("==========迷宫地图==========\n")
	for i := 0; i < len(mazeMap); i++ {

		for j := 0; j < len(mazeMap[i]); j++ {
			fmt.Print(mazeMap[i][j], " ")
		}

		fmt.Println()

	}

	GetMazeWay(&mazeMap, 1, 1)
	fmt.Println()
	fmt.Println("==========迷宫地图&走迷宫的路径==========\n")
	for i := 0; i < len(mazeMap); i++ {

		for j := 0; j < len(mazeMap[i]); j++ {
			fmt.Print(mazeMap[i][j], " ")
		}

		fmt.Println()

	}

}
