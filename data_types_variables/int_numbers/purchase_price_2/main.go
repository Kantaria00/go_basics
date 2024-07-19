package main

import (
	"fmt"
)

func main() {
	var a, b, n int
	_, err := fmt.Scan(&a, &b, &n)
	if err != nil {
		return
	}
	a = a*n + b*n/100
	b = b * n % 100
	fmt.Println(a, b)
}
