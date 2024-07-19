package main

import (
	"fmt"
)

func main() {
	var num int16
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Println(num % 1000 / 100)
}
