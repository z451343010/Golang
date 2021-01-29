/*
	测试生成订单号
*/
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/xingliuhua/leaf"
)

func main() {

	err, node := leaf.NewNode(0)
	if err != nil {
		return
	}
	// 每毫秒200个
	err = node.SetGenerateIDRate(200)
	if err != nil {
		return
	}
	startTime := time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	// 从2021年开始
	err = node.SetSince(startTime)
	if err != nil {
		return
	}
	// for i := 0; i < 400; i++ {
	// 	err, id := node.NextId()
	// 	if err != nil {
	// 		return
	// 	}
	// 	fmt.Println(id)
	// }

	err, id := node.NextId()
	if err != nil {
		return
	}

	id = strings.ToUpper(id)

	str := "028m2owm000"
	str = strings.ToUpper(str)

	fmt.Println(id)
	fmt.Println("str = ", str)

}
