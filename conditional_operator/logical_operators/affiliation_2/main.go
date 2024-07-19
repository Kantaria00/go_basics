package main

import (
	"fmt"
)

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	if a <= -3 || a >= 7 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}
