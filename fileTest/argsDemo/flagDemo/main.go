package main

import (
	"flag"
	"fmt"
)

func main() {

	// 定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	// StringVar() 参数说明
	// "u", -u 指定参数
	// "",默认值
	// "用户名，默认为空",说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	// func Parse
	// 必须在所有flag都注册好而未访问其值时执行
	// 未注册却使用 flag-help 时，会返回 ErrHelp
	flag.Parse()

	// 输出结果
	fmt.Printf("user = %v pwd = %v host = %v port = %v",
		user, pwd, host, port)

}
