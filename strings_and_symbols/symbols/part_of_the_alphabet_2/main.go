package main

import (
	"fmt"
)

func main() {
	var c string
	_, _ = fmt.Scan(&c)
	r := []rune(c)[0]
	for i := r; i <= 'z'; i++ {
		fmt.Print(string(i), " ")
	}
}
