package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	_, err := fmt.Scan(&x1, &y1, &x2, &y2)
	if err != nil {
		return
	}
	res := math.Abs(x1-x2) + math.Abs(y1-y2)
	fmt.Print(res)
}
