package main

import (
	"fmt"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)
	for N != 0 {
		fmt.Print(N % 2)
		N /= 2
	}
}
