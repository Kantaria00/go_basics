package main

import (
	"fmt"
)

func main() {
	var n, h, res int
	_, _ = fmt.Scan(&n, &h)
	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
		if arr[i] > h {
			res += 2
		} else {
			res++
		}
	}
	fmt.Println(res)
}
