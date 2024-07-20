package main

import (
	"fmt"
)

func main() {
	var n, m int
	flag := "#Black&White"
	_, _ = fmt.Scan(&n, &m)
	arr := make([][]string, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]string, m)
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&arr[i][j])
			if arr[i][j] != "W" && arr[i][j] != "B" && arr[i][j] != "G" {
				flag = "#Color"
				break
			}
		}
	}
	fmt.Println(flag)
}
