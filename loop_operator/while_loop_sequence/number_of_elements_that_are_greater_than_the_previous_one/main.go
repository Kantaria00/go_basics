package main

import "fmt"

func main() {
	var n, cnt, a int
	_, _ = fmt.Scan(&n)
	a = n
	for n != 0 {
		_, _ = fmt.Scan(&n)
		if a < n {
			cnt++
		}
		a = n
	}
	fmt.Println(cnt)
}
