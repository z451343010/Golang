/*
	Golang 操作 Redis
*/
package main

import (
	"fmt"
	// 引入 Redis 包
	"github.com/garyburd/redigo/redis"
)

func main() {

	// 链接到 Redis
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println("connect redis fails err = ", err)
		return
	}

	fmt.Println(conn)
	fmt.Println("connect success")

	// 通过 go 向 Redis 写入数据
	_, err = conn.Do("Set", "name", "tom猫")
	if err != nil {
		fmt.Println("set redis err = ", err)
		return
	}

	// 通过 go 向 Redis 读取数据
	result, readErr := conn.Do("Get", "name")
	if readErr != nil {
		fmt.Println("get Redis readErr = ", readErr)
		return
	}

	// 因为返回的 result 是 interface{} 类型
	// 因此我们需要把它转换成相应的类型，使用 redis 的 String 方法
	// String func (reply interface{},err error) (string,error)
	resultStr, transErr := redis.String(result, readErr)
	if transErr != nil {
		fmt.Println("trans Redis transErr = ", transErr)
		return
	}

	fmt.Println(resultStr)

	// 关闭链接
	defer conn.Close()

}
