package main

import (
	"fmt"
)

func main() {
	var n int
	count := 1
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
		if i != 0 && arr[i] != arr[i-1] {
			count++
		}
	}
	fmt.Println(count)
}
