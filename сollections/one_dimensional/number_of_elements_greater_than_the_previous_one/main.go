package main

import "fmt"

func main() {
	var n, cnt int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	_, _ = fmt.Scan(&n)
	arr[0] = n
	for i := 1; i < len(arr); i++ {
		_, _ = fmt.Scan(&n)
		arr[i] = n
		if arr[i] > arr[i-1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
