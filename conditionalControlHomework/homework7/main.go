package main

import (
	"fmt"
	"math"
)

// 根据公式：（身高 - 108）* 2 = 体重
// 可以有十斤左右的浮动，来观察测试者体重是否合格
func main() {

	var height float64
	var weight float64

	fmt.Println("请输入身高：")
	fmt.Scanln(&height)
	fmt.Println("请输入体重：")
	fmt.Scanln(&weight)

	if math.Abs((height-108)*2-weight) <= 10 {
		fmt.Println("测试者体重合格")
	} else {
		fmt.Println("测试者体重不合格")
	}

}
