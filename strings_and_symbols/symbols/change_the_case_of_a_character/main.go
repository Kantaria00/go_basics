package main

import (
	"fmt"
)

func main() {
	var c string
	_, _ = fmt.Scan(&c)
	r := []rune(c)[0]
	if r >= 'a' && r <= 'z' {
		fmt.Println(string(r - 'a' + 'A'))
	} else {
		fmt.Println(string(r - 'A' + 'a'))
	}
}
