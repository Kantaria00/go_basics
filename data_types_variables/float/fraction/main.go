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
	fmt.Print(num - float64(int(num)))
}
