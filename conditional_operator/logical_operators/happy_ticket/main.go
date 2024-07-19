package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a)
	b = a / 1000
	a %= 1000
	if (a/100 + a%100/10 + a%10) == (b/100 + b%100/10 + b%10) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
