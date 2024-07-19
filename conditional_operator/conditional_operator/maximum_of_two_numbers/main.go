package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	_, _ = fmt.Scan(&num1, &num2)
	fmt.Println(math.Max(num1, num2))
}
