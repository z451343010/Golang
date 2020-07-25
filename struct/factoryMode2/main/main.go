package main

import (
	"fmt"
	"struct/factoryMode2/model"
)

func main() {

	student := model.NewStudent("zjune", 100) // 指针类型的变量
	fmt.Println(*student)

	fmt.Println((*student).GetScore())

}
