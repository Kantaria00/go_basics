package main

import (
	"fmt"
)

func main() {
	var N, a, sum int
	_, _ = fmt.Scan(&N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&a)
		sum += a
	}
	fmt.Println(sum)
}
