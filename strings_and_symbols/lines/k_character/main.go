package main

import (
	"fmt"
)

func main() {
	var str, res string
	var k int
	_, _ = fmt.Scan(&str, &k)
	for i, a := range str {
		if i+1 != k {
			res += string(a)
		}
	}
	fmt.Println(res)
}
