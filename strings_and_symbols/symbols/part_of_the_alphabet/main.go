package main

import (
	"fmt"
)

func main() {
	var c string
	_, _ = fmt.Scan(&c)
	r := []rune(c)[0]
	for i := 97; i <= int(r); i++ {
		fmt.Print(string(rune(i)), " ")
	}
}
