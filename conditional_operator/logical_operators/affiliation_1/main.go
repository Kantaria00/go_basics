package main

import (
	"fmt"
)

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	if a > -1 && a < 17 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}
