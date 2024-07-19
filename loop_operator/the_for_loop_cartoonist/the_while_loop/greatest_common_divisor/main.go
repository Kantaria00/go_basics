package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	if b > a {
		temp := a
		a = b
		b = temp
	}
	for a%b != 0 {
		temp1 := b
		b = a % b
		a = temp1
	}
	fmt.Print(b)
}
