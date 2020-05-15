package main

import "fmt"

// 判断求婚者是否为高富帅
func main() {

	var height float64
	var treasure float64
	var isHandsome bool

	fmt.Println("请输入身高（厘米）：")
	fmt.Scanln(&height)
	fmt.Println("请输入财产(万)：")
	fmt.Scanln(&treasure)
	fmt.Println("请输入是否帅（true/false）：")
	fmt.Scanln(&isHandsome)

	if height >= 180 && treasure >= 1000 &&
		isHandsome == true {
		fmt.Println("我一定要嫁给他")
	} else if height >= 180 || treasure >= 1000 ||
		isHandsome == true {
		fmt.Println("比上不足，比下有余，嫁吧")
	} else {
		fmt.Println("不嫁")
	}

}
