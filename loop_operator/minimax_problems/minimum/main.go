package main

import "fmt"

func main() {
	var n, num int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&num)
	min := num
	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&num)
		if num < min {
			min = num
		}
	}
	fmt.Println(min)
}
