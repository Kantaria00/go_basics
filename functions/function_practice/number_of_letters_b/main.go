package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Println(sumB(s1) + sumB(s2))
}
func sumB(s string) int {
	var count int
	_, _ = fmt.Scan(&s)
	for _, i := range s {
		if i == 'b' {
			count++
		}
	}
	return count
}
