package main

import (
	"fmt"
)

func main() {
	var a, b int16
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}
	fmt.Print(float64(a+b) / 2.0)
}
