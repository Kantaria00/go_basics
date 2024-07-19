package main

import (
	"fmt"
)

func main() {
	var N, a, sum int
	_, _ = fmt.Scan(&N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&a)
		if a%2 == 0 && a%3 != 0 {
			sum += a
		}
	}
	fmt.Println(sum)
}
