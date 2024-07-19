package main

import (
	"fmt"
)

func main() {
	var N, count int
	_, _ = fmt.Scan(&N)
	for N != 0 {
		if N%2 == 1 {
			count++
		}
		N /= 2
	}
	fmt.Println(count)
}
