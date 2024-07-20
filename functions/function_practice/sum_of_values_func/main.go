package main

import (
	"fmt"
	"math"
)

func main() {
	var sum int
	for x := -25; x <= 15; x++ {
		sum += findY(x)
	}
	fmt.Println(sum)
}
func findY(x int) int {
	switch {
	case x < (-5):
		return 2*int(math.Abs(float64(x))) - 1
	case x > 5:
		return 2 * x
	default:
		return x
	}
}
