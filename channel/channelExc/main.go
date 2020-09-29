package main

import "fmt"

/*
	看看下面的代码会输出什么
*/

type Cat struct {
	Name string
	Age  int
}

func main() {

	var allChan chan interface{}
	allChan = make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat

	// 要得到第三个元素，需要把前2个元素出队列
	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat type = %T , newCat = %v\n", newCat, newCat)
	theCat := newCat.(Cat)
	fmt.Println("newCat.Name = ", theCat.Name)

}
