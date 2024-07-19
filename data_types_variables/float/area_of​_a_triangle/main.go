package main

import (
	"fmt"
)

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}
	fmt.Print(a * b / 2.0)
}
