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
		if arr[i]%10 == 7 && arr[i]%3 == 0 {
			count++
		}
	}
	for i := 0; i < n; i++ {
		if arr[i]%10 == 7 && arr[i]%3 == 0 {
			arr[i] = count
		}
		fmt.Print(arr[i], " ")
	}
}
