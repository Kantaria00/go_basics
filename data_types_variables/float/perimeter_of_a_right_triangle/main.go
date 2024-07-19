package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}
	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Print(c + a + b)
}
