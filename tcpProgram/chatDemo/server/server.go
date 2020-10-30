package main

import (
	"fmt"
	"net"
)

// 和客户端进行交互的协程
func process(conn net.Conn) {

	defer conn.Close()

	// 循环接收客户端发送的数据
	for {

		buf := make([]byte, 1024)
		// 1.等待客户端通过conn发送信息
		// 2.如果客户端没有write[发送]，那么协程就阻塞在这里
		fmt.Println("服务器等待客户端输入......", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn读取
		if err != nil {
			fmt.Println("客户端已经退出")
			return
		}

		var message string = string(buf[:(n - 2)])

		if message == "exit" {
			fmt.Println("客户端请求退出")
			return
		}

		// 3.显示客户端发送的内容到服务器的终端
		fmt.Println("客户端输入的内容为：")
		fmt.Printf("%v:%v\n\n", conn.RemoteAddr().String(), message)
	}

}

func main() {

	fmt.Println("服务器开始监听")
	// 在本地监听 8888 端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}

	// 关闭监听资源
	defer listen.Close()

	// 循环等待客户端来链接
	for {
		fmt.Println()
		fmt.Println("等待客户端来连接......")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("conn err = ", err)
			return
		} else {
			fmt.Println("conn success conn = ", conn)
			fmt.Println("client address = ", conn.RemoteAddr().String())
		}

		// 准备一个协程，为客户端服务
		go process(conn)

	}

	// fmt.Println("listen = ", listen)

}
