package main

import "fmt"

/*
	根据行、列、字符打印对应行数和列数的字符
	比如：行：3，列：2,字符*,则打印相应的结果
*/

type Test struct {
}

func (test *Test) printChar(multiArr [3][3]byte, row int,
	line int) { // 行，列

	fmt.Printf("%c\n", multiArr[row-1][line-1])

}

func main() {

	var multiArr [3][3]byte = [3][3]byte{{'x', 'y', 'z'},
		{'a', 'b', 'c'}, {'&', '*', '('}}
	var row int = 3
	var line int = 2

	var test Test
	(&test).printChar(multiArr, row, line)

}
