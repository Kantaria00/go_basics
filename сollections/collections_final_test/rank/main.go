package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := n + 1
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	for i := n; i > 0; i-- {
		if arr[i] > arr[i-1] {
			tmp := arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp
			num--
		} else {
			break
		}
	}
	fmt.Println(num)
}
