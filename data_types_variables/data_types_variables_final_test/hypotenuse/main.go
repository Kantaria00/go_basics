package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		return
	}
	fmt.Println(math.Sqrt(math.Pow(num1, 2) + math.Pow(num2, 2)))
}
