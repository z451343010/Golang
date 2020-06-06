package main

import (
	"fmt"
	"time"
)

// 休眠函数
func main() {

	// 每隔1秒钟打印一个数字，打印到5的时候就退出
	var i int = 0
	for {

		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Second)

		if i == 5 {
			break
		}

	}

	timeNow := time.Now()
	// unix和unixnano的使用
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n",
		timeNow.Unix(), timeNow.UnixNano())

}
