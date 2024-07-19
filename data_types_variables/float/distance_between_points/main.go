package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 int
	_, err := fmt.Scan(&x1, &y1, &x2, &y2)
	if err != nil {
		return
	}
	res := math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
	fmt.Print(res)
}
