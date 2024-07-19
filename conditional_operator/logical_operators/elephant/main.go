package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	_, _ = fmt.Scan(&x1, &x2, &y1, &y2)
	if math.Abs(x1-y1) == math.Abs(x2-y2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
