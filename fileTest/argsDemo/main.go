package main

import (
	"fmt"
	"os"
)

/*
	编写一段代码，获取各个命令行的参数
*/
func main() {

	fmt.Println("命令行的参数有：", len(os.Args), "个")

	for index, value := range os.Args {
		fmt.Printf("args[%v] = %v\n", index, value)
	}

}
