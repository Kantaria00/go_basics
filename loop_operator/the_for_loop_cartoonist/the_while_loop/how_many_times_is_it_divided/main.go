package main

import (
	"fmt"
)

func main() {
	var n, count int
	_, _ = fmt.Scan(&n)
	a := 3
	for n%a == 0 {
		a *= 3
		count++
	}
	fmt.Println(count)
}
