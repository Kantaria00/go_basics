package main

import (
	"fmt"
)

func main() {
	var n int
	res := 1
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			res *= i
		}
	}
	fmt.Println(res)
}
