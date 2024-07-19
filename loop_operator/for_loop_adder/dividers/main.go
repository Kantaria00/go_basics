package main

import (
	"fmt"
)

func main() {
	var x, count int
	_, _ = fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			count++
		}
	}
	fmt.Println(count)
}
