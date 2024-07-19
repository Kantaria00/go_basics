package main

import (
	"fmt"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)
	for N != 0 {
		fmt.Print(N % 10)
		N /= 10
	}
}
