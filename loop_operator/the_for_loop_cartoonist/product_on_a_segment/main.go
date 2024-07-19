package main

import (
	"fmt"
)

func main() {
	var a, b int
	res := 1
	_, _ = fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		res *= i
	}
	fmt.Println(res)
}
