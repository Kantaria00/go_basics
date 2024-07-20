package main

import (
	"fmt"
)

func main() {
	var n, num int
	cnt := 1
	_, _ = fmt.Scan(&n, &num)
	mn := num
	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&num)
		switch {
		case num == mn:
			cnt++
		case num < mn:
			mn = num
			cnt = 1
		}
	}
	fmt.Println(cnt)
}
