package main

import (
	"fmt"
)

func main() {
	var a, res int32
	for i := 0; i < 3; i++ {
		_, _ = fmt.Scan(&a)
		if a%2 == 0 {
			res += a / 2
		} else {
			res += a/2 + 1
		}
	}
	fmt.Println(res)
}
