package main

import (
	"fmt"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		fmt.Println(i * i)
	}
}
