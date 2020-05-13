package main

import "fmt"

func main() {

	var num1 int32 = 25
	var num2 int32 = 50

	if (num1+num2)%3 == 0 {
		if (num1+num2)%5 == 0 {
			fmt.Println("这个数既能够被3整除，又能够被5整除")
		}
	}

}
