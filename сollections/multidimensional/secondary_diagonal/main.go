package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			switch {
			case i+j == n-1:
				arr[i][j] = 1
			case i+j < n-1:
				arr[i][j] = 0
			case i+j > n-1:
				arr[i][j] = 2
			}
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
