package main

import "fmt"

/*
	编程创建一个Box结构体，在其中声明三个字段
	表示一个立方体的长、宽、高
	长宽高要从终端获取
	声明获取一个立方体的体积的方法
	创建一个Box结构体变量，打印给定尺寸的立方体的体积
*/

type Box struct {
	long   float64
	width  float64
	height float64
}

func (box *Box) getVolume() (volume float64) {

	volume = (*box).height * (*box).width * (*box).long
	return volume

}

func main() {

	var box Box
	var long float64
	var width float64
	var height float64

	fmt.Scanln(&long)
	fmt.Scanln(&width)
	fmt.Scanln(&height)

	box.height = height
	box.long = long
	box.width = width

	fmt.Println((&box).getVolume())

}
