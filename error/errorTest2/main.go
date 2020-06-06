package main

import (
	"errors"
	"fmt"
)

// 函数去读取配置文件conf.ini里面的信息
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(fileName string) (err error) {

	if fileName == "conf.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}

}

func test2() {

	err := readConf("config.ini")
	if err != nil {
		// 如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}

	// 判断函数是否继续执行
	fmt.Println("test2()继续执行")

}

func main() {
	test2()
}
