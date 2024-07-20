package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	a := strconv.Itoa(n)
	for _, i := range a {
		if i != 53 && i != 55 {
			fmt.Print(string(i))
		}
	}
}
