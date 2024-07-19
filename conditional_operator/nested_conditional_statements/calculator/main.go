package main

import (
	"fmt"
)

func main() {
	var x, y int
	var z string
	_, _ = fmt.Scan(&x, &y, &z)
	switch {
	case z == "+":
		fmt.Println(x + y)
	case z == "-":
		fmt.Println(x - y)
	case z == "*":
		fmt.Println(x * y)
	case z == "/":
		if y == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println(float64(x) / float64(y))
		}
	default:
		fmt.Println("Неверная операция")
	}
}
