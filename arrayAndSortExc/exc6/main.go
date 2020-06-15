package main

import "fmt"

// 已知数组 arr[10] string 里面保存了10个元素
// 查找AA
// 并且打印所有AA的坐标
func main() {

	var strArr [10]string

	fmt.Println("请依次输入10个字符串：")

	for i := 0; i < len(strArr); i++ {
		var str string
		fmt.Scanln(&str)
		strArr[i] = str
	}

	fmt.Println("带有\"AA\"的数组下标为：")

	for i := 0; i < len(strArr); i++ {

		var charSlice []byte

		for j := 0; j < len(strArr[i]); j++ {
			charSlice = append(charSlice, byte(strArr[i][j]))
		}

		for k := 0; k < len(charSlice); k++ {

			if k < len(charSlice)-1 &&
				charSlice[k] == 'A' && charSlice[k+1] == 'A' {

				fmt.Print(i, " ")
				break

			}
		}

	}

}
