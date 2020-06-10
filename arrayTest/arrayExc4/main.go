package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成5个数，并将其反转打印
func main() {

	var randNumArr [5]int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(randNumArr); i++ {
		var randNum int = rand.Intn(100)
		randNumArr[i] = randNum
	}

	fmt.Println(randNumArr)

	// 直接反转打印
	// for j := (len(randNumArr) - 1); j >= 0; j-- {
	// 	fmt.Println(randNumArr[j])
	// }

	// 直接反转数组内容
	for k := 0; k < len(randNumArr)/2; k++ {
		var temp int
		temp = randNumArr[len(randNumArr)-1-k]
		randNumArr[len(randNumArr)-1-k] = randNumArr[k]
		randNumArr[k] = temp

	}

	fmt.Println("反转后的数组为：")
	fmt.Println(randNumArr)

}
