package main

import "fmt"

/*
	多态
*/

// 声明/定义一个接口
type Usb interface {

	// 声明了两个没有实现的方法
	Start()
	Stop()
}

// 结构体 Phone
type Phone struct {
	Name string
}

// Phone 实现 Usb接口的方法
func (phone *Phone) Start() {
	fmt.Println("Phone start")
}

func (phone *Phone) Stop() {
	fmt.Println("Phone stop")
}

// 结构体 Camera
type Camera struct {
	Name string
}

// Camera 实现 Usb接口的方法
func (camera *Camera) Start() {
	fmt.Println("Camera start")
}

func (camera *Camera) Stop() {
	fmt.Println("Camera stop")
}

func main() {

	// 定义个Usb接口数组，可以存放 Phone 和 Camera 的结构体变量
	var usbArr [3]Usb
	usbArr[0] = &Phone{"HUAWEI"}
	usbArr[1] = &Phone{"XIAOMI"}
	usbArr[2] = &Camera{"CANON"}

}
