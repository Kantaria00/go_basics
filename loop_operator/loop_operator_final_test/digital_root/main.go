package main

import (
	"fmt"
)

func main() {
	var n, res int
	_, _ = fmt.Scan(&n)
	if n < 10 {
		res = n
	} else {
		for n/10 != 0 {
			res = 0
			for n != 0 {
				res += n % 10
				n /= 10
			}
			n = res
		}
	}

	fmt.Println(res)
}
