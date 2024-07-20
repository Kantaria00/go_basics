package main

import (
	"fmt"
)

func main() {
	var a, n int
	res := 1
	_, _ = fmt.Scan(&a, &n)
	for n != 0 {
		res *= a
		n--
	}
	fmt.Println(res)
}
