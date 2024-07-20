package main

import (
	"fmt"
)

func main() {
	var a1, a2 string
	_, _ = fmt.Scan(&a1, &a2)
	r1, r2 := []rune(a1)[0], []rune(a2)[0]
	if r1 < r2 {
		for i := r1; i <= r2; i++ {
			fmt.Print(string(i), " ")
		}
	} else {
		for i := r2; i <= r1; i++ {
			fmt.Print(string(i), " ")
		}
	}
}
