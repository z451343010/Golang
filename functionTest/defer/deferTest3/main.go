package main

func test() {
	// 关闭文件资源
	file = openfile(文件名)
	defer file.close()
	// 其它代码
}

func test2() {
	// 释放数据库连接资源
	connect = openDatabase()
	defer connect.close()
	// 其它代码
}
