package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Print(num / math.Pow(2, 13))
}
