package main

import (
	"fmt"
)

func main() {
	var x float64
	_, _ = fmt.Scan(&x)
	if x < 60 {
		fmt.Println("Легкий вес")
	} else if x < 64 {
		fmt.Println("Первый полусредний вес")
	} else if x < 69 {
		fmt.Println("Полусредний вес")
	}
}
