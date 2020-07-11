package main

import "fmt"

// 声明一个结构体Circle，字段为radius
// 声明一个方法getArea()和Circle进行绑定，可以返回面积

type Circle struct {
	Radius float64
}

func (circle Circle) getArea(radius float64) (area float64) {

	area = radius * radius * 3.14
	return area

}

// 为了提高效率，通常让方法与结构体的指针类型绑定
func (circle *Circle) getArea2(radius float64) (area float64) {

	fmt.Printf("getArea2 circle2 %p\n", circle)
	area = radius * radius * 3.14
	return area

}

func main() {

	var circle Circle
	var radius float64
	radius = 5
	fmt.Println("the are of circle is:", circle.getArea(radius))

	var circle2 Circle
	circle2.Radius = 5.0
	fmt.Printf("main address circle2 %p\n", &circle2)
	result2 := (&circle2).getArea2(circle2.Radius)
	fmt.Println("circle2 ", result2)

}
