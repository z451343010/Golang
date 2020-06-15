package main

import "fmt"

func main() {

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("地址arr2[0]:%p\n", &arr2[0])
	fmt.Printf("地址arr2[1]:%p\n", &arr2[1])
	fmt.Printf("地址arr2[0][1]:%p", &arr2[0][1])
}
