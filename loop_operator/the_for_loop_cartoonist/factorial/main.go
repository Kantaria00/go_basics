package main

import (
	"fmt"
)

func main() {
	var N int
	var res int = 1
	_, _ = fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		res *= i
	}
	fmt.Println(res)
}
