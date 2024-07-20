package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			switch {
			case i == 0 || j == 0:
				arr[i][j] = 1
			default:
				arr[i][j] = arr[i-1][j] + arr[i][j-1]
			}
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
