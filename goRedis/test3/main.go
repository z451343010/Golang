/*
	批量 Set/Get 数据
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

	_, err = conn.Do("MSet", "name", "马士兵", "address", "北京")
	if err != nil {
		fmt.Println("MSet err = ", err)
		return
	}

	result, err := redis.Strings(conn.Do("MGet", "name", "address"))
	if err != nil {
		fmt.Println("Mget err = ", err)
		return
	}

	for _, value := range result {
		fmt.Println(value)
	}

	// HMSet HMGet
	_, err = conn.Do("HMSet", "user2", "name", "John", "age", 19)
	if err != nil {
		fmt.Println("HMset err = ", err)
		return
	}

	hmResult, err := redis.Strings(conn.Do("HMGet", "user2", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err = ", err)
	}

	for _, value := range hmResult {
		fmt.Println(value)
	}

}
