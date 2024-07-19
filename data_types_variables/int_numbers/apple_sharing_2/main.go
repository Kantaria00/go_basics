package main

import (
	"fmt"
)

func main() {
	var n, k int16
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	fmt.Print(k % n)
}
