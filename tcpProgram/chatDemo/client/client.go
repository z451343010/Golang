package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func sendMessage(conn net.Conn) {

	// 功能一：客户端可以发送单行数据，然后就退出
	// os.Stdin 代表标准输入（从终端的键盘读取输入）
	reader := bufio.NewReader(os.Stdin)

	for {

		// 读到回车（换行）结束
		line, readErr := reader.ReadString('\n')
		if readErr != nil {
			fmt.Println("readErr = ", readErr)
			return
		}

		// 将读取到的字符串发送给服务器
		n, cwErr := conn.Write([]byte(line))
		if cwErr != nil {
			fmt.Println("conn.Write err = ", cwErr)
		}

		fmt.Printf("客户端发送了 %d byte 的数据，并退出\n", n)

	}

}

func main() {

	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("conn fail conn err = ", err)
		return
	} else {
		fmt.Println("conn success conn = ", conn)
	}

	// // 功能一：客户端可以发送单行数据，然后就退出
	// // os.Stdin 代表标准输入（从终端的键盘读取输入）
	// reader := bufio.NewReader(os.Stdin)
	// // 读到回车（换行）结束
	// line, readErr := reader.ReadString('\n')
	// if readErr != nil {
	// 	fmt.Println("readErr = ", readErr)
	// 	return
	// }

	// // 将读取到的字符串发送给服务器
	// n, cwErr := conn.Write([]byte(line))
	// if cwErr != nil {
	// 	fmt.Println("conn.Write err = ", cwErr)
	// }

	// fmt.Printf("客户端发送了 %d byte 的数据，并退出\n", n)

	for {

		sendMessage(conn)

	}

}
