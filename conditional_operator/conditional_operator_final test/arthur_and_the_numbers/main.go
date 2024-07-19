package main

import (
	"fmt"
	"math"
)

func main() {
	var k2, k3, k5, k6 float64
	_, _ = fmt.Scan(&k2, &k3, &k5, &k6)
	x256 := math.Min(k2, math.Min(k5, k6))
	x32 := math.Min(k2-x256, k3)
	fmt.Println(256*x256 + 32*x32)
}
