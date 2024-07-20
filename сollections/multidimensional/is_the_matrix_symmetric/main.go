package main

import (
	"fmt"
)

func main() {
	var n int
	flag := "YES"
	_, _ = fmt.Scan(&n)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&arr[i][j])
			if i >= j {
				if arr[i][j] != arr[j][i] {
					flag = "NO"
					break
				}
			}
		}
		if flag == "NO" {
			break
		}

	}
	fmt.Println(flag)
}
