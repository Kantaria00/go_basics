package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	_, _ = fmt.Scan(&x, &y, &z)
	if x == y && x == z && y == z {
		fmt.Println(3)
	} else if x != y && x != z && y != z {
		fmt.Println(0)
	} else {
		fmt.Println(2)
	}
}
