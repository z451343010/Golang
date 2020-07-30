package main

import (
	"fmt"
	"struct/packageExc/model"
)

func main() {

	account := model.NewAccount("z451343010")
	(*account).SetPassword("123456123")
	(*account).SetBalance(125)

	fmt.Println(*account)

}
