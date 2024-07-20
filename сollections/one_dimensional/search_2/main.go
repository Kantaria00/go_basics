package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
		if arr[i]%3 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
}
