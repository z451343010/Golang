package main

import "fmt"

// 定义一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

type Camera struct {
}

// 让Phone实现Usb接口的方法
func (p *Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p *Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让Camera实现Usb接口的方法
func (p *Camera) Start() {
	fmt.Println("相机开始工作")
}

func (p *Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

// 只要是实现了Usb接口（所谓实现Usb接口，
// 就是指实现了Usb接口声明的所有方法）
func (c *Computer) Working(usb Usb) {

	// 通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()

}

func main() {

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(&phone)
	computer.Working(&camera)

}
