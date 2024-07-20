package main

import (
	"fmt"
)

func main() {
	var n int
	flag := "NO"
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				flag = "YES"
				break
			}
		}
		if flag == "YES" {
			break
		}
	}
	fmt.Println(flag)
}
