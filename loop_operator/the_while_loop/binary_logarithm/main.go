package main

import (
	"fmt"
)

func main() {
	var n, st int
	_, _ = fmt.Scan(&n)
	a := 1

	for a < n {
		a *= 2
		st++
	}
	fmt.Println(st)
}
