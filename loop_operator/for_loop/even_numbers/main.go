package main

import (
	"fmt"
)

func main() {
	var a, b int32
	_, _ = fmt.Scan(&a, &b)
	if a%2 != 0 {
		a++
	}
	for i := a; i <= b; i += 2 {
		fmt.Println(i)
	}
}
