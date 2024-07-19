package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	_, err := fmt.Scan(&num1)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&num2)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&num3)
	if err != nil {
		return
	}
	fmt.Print(num1 * num2 * num3)
}
