package main

import "fmt"

func main() {
	var n, ind, mn int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
		if i == 0 || arr[i] < mn {
			mn = arr[i]
			ind = i
		}
	}
	fmt.Println(ind)
}
