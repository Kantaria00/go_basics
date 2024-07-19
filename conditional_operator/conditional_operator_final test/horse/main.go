package main

import (
	"fmt"
)

func main() {
	var x1, x2, y1, y2 int8
	_, _ = fmt.Scan(&x1, &x2, &y1, &y2)
	switch {
	case (x1+1 == y1) && (x2+2 == y2) ||
		(x1+2 == y1) && (x2+1 == y2) ||
		(x1-1 == y1) && (x2+2 == y2) ||
		(x1-2 == y1) && (x2+1 == y2) ||
		(x1-1 == y1) && (x2-2 == y2) ||
		(x1-2 == y1) && (x2-1 == y2) ||
		(x1+1 == y1) && (x2-2 == y2) ||
		(x1+2 == y1) && (x2-1 == y2):
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
