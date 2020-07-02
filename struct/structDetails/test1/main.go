package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftPoint  Point
	rightPoint Point
}

// 结构体的使用细节
// 结构体的所有字段在内存中是连续的
func main() {

	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println("r1.leftPoint.x 的地址：", &r1.leftPoint.x)
	fmt.Println("r1.leftPoint.y 的地址：", &r1.leftPoint.y)
	fmt.Println("r1.rightPoint.x 的地址：", &r1.rightPoint.x)
	fmt.Println("r1.rightPoint.x 的地址：", &r1.rightPoint.y)

}
