package main

import (
	"fmt"
)

func main() {
	var n, count int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
		if arr[i]%2 == 0 {
			count++
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
	fmt.Println(count)
}
