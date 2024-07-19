package main

import (
	"fmt"
)

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	num1 := num * num
	num2 := num1 * num1
	num3 := num1 * num2
	fmt.Print(num3)
}
