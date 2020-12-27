/*
	数据结构
	队列
	题目一：
	1）创建一个数组模拟队列，每隔一定时间[随机]，给数组添加一个数。
	2）启动两个协程，每隔一定时间（时间随机）到队列取出数据
	3）在控制台输出
		x 号协程 服务 ---》x 号客户
		x 号协程 服务 ---》x 号客户
		x 号协程 服务 ---》x 号客户
	4）使用锁机制即可
*/
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	maxSize int
	queue   [15]int
	front   int
	rear    int
}

// 声明一个全局互斥锁
var lock sync.Mutex

// 添加数据到队列
func (this *Queue) AddQueue(num int) (err error) {

	// 判断队列是否已满
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}

	this.rear++
	this.queue[this.rear] = num
	return

}

// 每隔一定的随机时间
// 添加数据到队列
func (this *Queue) AddQueueRandomTime() {

	num := 0
	randNum := 0

	// 设置随机源
	rand.Seed(time.Now().UnixNano())

	for {
		// 程序睡眠随机一段时间
		randNum = rand.Intn(5000)
		time.Sleep(time.Millisecond * time.Duration(randNum))
		num++
		this.AddQueue(num)

		if this.rear == this.maxSize-1 {
			break
		}

	}

}

// 出队列
func (this *Queue) GetQueue() (value int, err error) {

	// 先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	value = this.queue[this.front]
	return value, err
}

// 每隔随机一段时间
// 从队列中取出数据
func (this *Queue) GetQueueRandomTime(routineNum int) {

	// 设置随机源
	rand.Seed(time.Now().UnixNano())

	for {

		// 程序睡眠随机一段时间
		randNum := rand.Intn(5000)
		time.Sleep(time.Millisecond * time.Duration(randNum))

		// 锁住队列
		// 当某个协程访问完这个队列
		// 另一个协程才允许访问这个队列
		lock.Lock()
		num, err := this.GetQueue()
		lock.Unlock()

		if err != nil {
			fmt.Println("队列为空")
		} else {
			fmt.Printf("%d 号协程 服务 ---》 %d 号客户\n", routineNum, num)
		}

	}

}

func main() {

	// 初始化队列
	queue := &Queue{
		maxSize: 15,
		front:   -1,
		rear:    -1,
	}

	// 开启往队列添加数据的协程
	go queue.AddQueueRandomTime()

	// 开启2个协程，往队列取出数据
	for i := 1; i <= 2; i++ {
		go queue.GetQueueRandomTime(i)
	}

	// 主线程进入睡眠
	// 防止程序直接结束
	time.Sleep(time.Second * 1000)

}
