package main

import (
	"fmt"
)

func main() {
	var N, del int
	_, _ = fmt.Scan(&N)
	a := N
	for a != 0 {
		del += a % 10
		a /= 10
	}
	if N%del == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
