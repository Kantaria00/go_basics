package main

import (
	"fmt"
)

func main() {
	var num float64
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Print(3.14 * num * num)
}
