package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 1; i < len(arr); i++ {
		_, _ = fmt.Scan(&n)
		arr[i] = n
	}
	_, _ = fmt.Scan(&n)
	arr[0] = n
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}
