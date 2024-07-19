package main

import (
	"fmt"
)

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	fmt.Print(k / n)
}
