/*
	Redis 的课堂练习1：
	1）Monster信息[name,age,skill]
	2）通过终端输入三个 monster 的信息，使用 Golang 操作 Redis
	3）编程，遍历出所有的 monter 信息，并显示在终端
	4）提示：可以使用 hash 数据类型
*/
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis err = ", err)
		return
	}

	_, err = conn.Do("HMSet", "monster", "name", "Diablo", "age", 3000,
		"skill", "Fire")
	if err != nil {
		fmt.Println("HMSet redis err = ", err)
		return
	}

	result, err := redis.Strings(conn.Do("HGetall", "monster"))
	if err != nil {
		fmt.Println("HGetall err = ", err)
		return
	}

	for _, value := range result {
		fmt.Println(value)
	}

}
