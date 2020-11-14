/*
	Golang redis 连接池
*/
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 定义一个全局变量 pool
var pool *redis.Pool

// 启动程序时，就初始化连接池
func init() {

	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 表示和数据库的最大连接，0表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}

func main() {

	// 先从 pool 取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "张zjune")
	if err != nil {
		fmt.Println("redis set err = ", err)
		return
	}

	result, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis get err = ", err)
		return
	}

	fmt.Println(result)

}
