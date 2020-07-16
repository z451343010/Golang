package main

import "fmt"

/*
	编写方法：判断一个数是奇数还是偶数
*/

type Judge struct {
}

func (judge *Judge) judgeNum(num int64) {

	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

}

func main() {

	var num int64
	num = 16
	var judge Judge
	(&judge).judgeNum(num)

}
