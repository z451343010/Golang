/*
	Redis 课堂练习2
	要求：
	1）记录用户浏览的信息，比如保存商品名
	2）编写一个函数，可以取出某个用户最近浏览的10个商品名
	3）提示：考虑使用 list 数据类型
*/
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect Redis err = ", err)
		return
	}

	_, err = conn.Do("LPush", "goodList", "Java", "PHP", "Kotlin",
		"C", "C++", "Go", "JavaScript", "HTML", "CSS", "Python",
		"Swift", "ObjectC", "Rust", "B")
	if err != nil {
		fmt.Println("lpush Redis err = ", err)
		return
	}

	goodNum, err := redis.Int(conn.Do("LLen", "goodList"))
	if err != nil {
		fmt.Println("LLen Redis err = ", err)
		return
	}

	if goodNum < 10 {
		fmt.Println("number of goods is lower than 10")
		return
	}

	for i := 0; i < 10; i++ {

		good, err := redis.String(conn.Do("LPop", "goodList"))
		if err != nil {
			fmt.Println("lpop Redis err = ", err)
			return
		}

		fmt.Println(good)

	}

}
