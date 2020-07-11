package main

import "fmt"

type Person struct {
}

func (person Person) calculate(n int) (sum int) {

	for i := 1; i <= n; i++ {

		sum += i

	}

	fmt.Println("calculate() n的地址：", &n)
	return sum

}

func main() {

	var person Person
	var n int
	n = 3
	fmt.Println("main() n 的地址：", &n)
	fmt.Println(person.calculate(n))

}
