package main

import "fmt"

// Go数组简单实例
func main() {

	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 2.4
	hens[4] = 3.5
	hens[5] = 50.0

	// 求总和
	var totalWeight float64
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	fmt.Println(totalWeight)

	// 求平均
	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Println(avgWeight)

}
