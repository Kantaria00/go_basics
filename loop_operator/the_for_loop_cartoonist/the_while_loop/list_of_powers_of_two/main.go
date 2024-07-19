package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	st := 1

	for st <= n {
		fmt.Println(st)
		st *= 2
	}
}
