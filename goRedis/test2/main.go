/*
	go Redis Hash
*/
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect fail err = ", err)
		return
	}

	_, hErr := conn.Do("Hset", "user1", "name", "John")
	if hErr != nil {
		fmt.Println("hash set Redis hErr = ", hErr)
		return
	}

	result, rErr := redis.String(conn.Do("Hget", "user1", "name"))
	if rErr != nil {
		fmt.Println("hash get Redis rErr = ", rErr)
		return
	}
	fmt.Println(result)

}
