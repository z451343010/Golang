package main

import "fmt"

func main() {

	var m int
	fmt.Scanln(&m)
	fmt.Println(m)

	for i := 1; i <= m; i++ {
		var v, p, q int
		fmt.Scanf("%d %d %d\n", &v, &p, &q)
	}

}
