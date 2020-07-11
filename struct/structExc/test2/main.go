package main

import "fmt"

// 给Person结构体添加 calculate() 方法
// 可以计算从1+...+1000的结果

type Person struct {
}

func (person Person) calculate() (sum int) {

	for i := 1; i <= 1000; i++ {
		sum += i
	}

	return sum

}

func main() {

	var person Person
	var sum int
	sum = person.calculate()
	fmt.Println("sum =", sum)

}
