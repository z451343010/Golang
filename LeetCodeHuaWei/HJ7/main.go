/*
	牛客网
	华为机试
	HJ7
	通过
*/
package main

import "fmt"

func main() {

	var num float64
	fmt.Scanln(&num)

	var temp int = int((num - float64(int(num))) * 10)

	if temp >= 5 {
		fmt.Print(int(num) + 1)
	} else {
		fmt.Print(int(num))
	}

}
