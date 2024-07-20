package main

import (
	"fmt"
)

func main() {
	var n, m, mx, ind int
	_, _ = fmt.Scan(&n, &m)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		sum := 0
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&arr[i][j])
			sum += arr[i][j]
		}
		if i == 0 || sum > mx {
			mx = sum
			ind = i
		}
	}
	fmt.Println(mx)
	fmt.Println(ind)
}
