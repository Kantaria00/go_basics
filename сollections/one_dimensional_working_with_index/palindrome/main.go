package main

import (
	"fmt"
)

func main() {
	var n int
	flag := "YES"
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[len(arr)-1-i] {
			flag = "NO"
			break
		}
	}
	fmt.Println(flag)
}
