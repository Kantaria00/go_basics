package main

import (
	"fmt"
)

func main() {
	var N, count int
	_, _ = fmt.Scan(&N)

	for N != 0 {
		if N%10 == 4 {
			count++
		}
		N /= 10
	}
	fmt.Println(count)
}
