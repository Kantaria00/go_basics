package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	res := 1
	_, _ = fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if math.Abs(float64(i%10)) == 7 {
			res *= i
		}
	}
	fmt.Println(res)
}
