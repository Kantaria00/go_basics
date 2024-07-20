package main

import "fmt"

func main() {
	var n, num, count int
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&num)
		if num == 0 {
			count++
		}
	}
	fmt.Println(count)
}
