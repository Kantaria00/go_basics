package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	_, _ = fmt.Scan(&a, &b, &c)
	D := math.Pow(b, 2) - 4*a*c
	switch {
	case D > 0:
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println(math.Min(x1, x2))
		fmt.Println(math.Max(x1, x2))
	case D == 0:
		x := -b / (2 * a)
		fmt.Println(x)
	}

}
