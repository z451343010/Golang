/*
	使用select可以解决从管道取数据的阻塞问题
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock

	// 问题，在实际开发中，可能我们不好确定什么时候关闭管道
	// 可以使用 select 方式解决
label:
	for {

		select {
		// 如果 intChan 一直没有关闭，不会一直阻塞而 deadlock
		// 会自动到下一个 case 匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("都取不到数据")
			time.Sleep(time.Second)
			break label

		}

	}

}
