package main

import (
	"fmt"
)

func main() {
	var a int8
	var s string
	_, _ = fmt.Scan(&a, &s)
	switch {
	case s == "m" && a >= 12 && a <= 18:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
