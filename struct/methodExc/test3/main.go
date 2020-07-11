package main

import "fmt"

/*
	编写一个方法，计算矩形的面积
	可以接收长 len ,宽 width
	将其作为方法返回值。在main中调用该方法
	接收返回的面积值，并打印
*/

type Rectangle struct {
}

func (rect *Rectangle) getArea(len float64, width float64) (area float64) {

	area = len * width
	return area

}

func main() {

	var rect Rectangle
	var len float64 = 3.14
	var width float64 = 3.15

	area := (&rect).getArea(len, width)

	fmt.Println(area)

}
