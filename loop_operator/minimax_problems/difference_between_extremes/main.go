package main

import "fmt"

func main() {
	var n, num int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&num)
	mx := num
	mn := num
	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&num)
		if num > mx {
			mx = num
		}
		if num < mn {
			mn = num
		}
	}
	fmt.Println(mx - mn)
}
