package main

import "fmt"

/*
	最佳实践案例
	在前面的Usb接口案例做改进：
	给Phone结构体增加一个特有的方法call()
	当Usb接口接收的是Phone变量时，还需要调用call方法
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

type Computer struct {
	Name string
}

type Camera struct {
	Name string
}

// Phone 实现 Usb接口的方法
func (phone Phone) Start() {
	fmt.Println(phone.Name, "Phone start")
}

func (phone Phone) Stop() {
	fmt.Println(phone.Name, "Phone stop")
}

func (phone Phone) Call() {
	fmt.Println(phone.Name, " phone call")
}

// Camera 实现 Usb接口的方法
func (camera Camera) Start() {
	fmt.Println(camera.Name, "Camera start")
}

func (camera Camera) Stop() {
	fmt.Println(camera.Name, "Camera stop")
}

func (computer Computer) Working(usb Usb) {

	usb.Start()

	// 如果Usb指向的是 Phone 结构体变量
	// 则还需要调用 Call() 方法
	// 类型断言
	phone, flag := usb.(Phone)

	// 判断接口类型
	if flag == true {

		phone.Call()

	} else {

		fmt.Println("is not belong Phone ")

	}

	usb.Stop()

}

func main() {

	var usbArr [3]Usb
	usbArr[0] = Phone{"Apple"}
	usbArr[1] = Camera{"Nikon"}
	usbArr[2] = Phone{"Huawei"}

	var computer Computer

	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}

}
