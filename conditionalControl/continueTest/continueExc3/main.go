package main

import "fmt"

// 某人有100,000元，每经过一次路口，需要交费，规则如下：
// 1.当现金>50000时，每次交5%
// 2.当现金<=50000时，每次交1000
func main() {

	var cash float64 = 100000
	var passNum float64

	for {

		if cash <= 0 {
			break
		}

		if cash > 50000 {
			cash -= cash * 0.05
			passNum++
			continue
		}

		cash -= 1000
		passNum++

	}

	fmt.Println("该人可经过", passNum, "次路口")

}
