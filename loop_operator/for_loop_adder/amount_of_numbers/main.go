package main

import (
	"fmt"
)

func main() {
	var N, a, count int
	_, _ = fmt.Scan(&N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&a)
		if a%10 == 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
