package main

import "fmt"

func main() {

	intChan := make(chan int, 50)
	intChan <- 8
	fmt.Println(len(intChan))

}
