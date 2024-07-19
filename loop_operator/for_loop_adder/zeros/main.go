package main

import (
	"fmt"
)

func main() {
	var N, a, count int
	_, _ = fmt.Scan(&N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&a)
		if a == 0 {
			count = 1
			break
		}
	}
	if count != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
