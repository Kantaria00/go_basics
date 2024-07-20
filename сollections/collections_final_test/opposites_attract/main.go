package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if arr[i]+arr[j] == 0 && i != j {
				fmt.Println(i, j)
				break
			}
		}
	}
}
