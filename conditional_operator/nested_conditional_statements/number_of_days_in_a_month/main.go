package main

import (
	"fmt"
)

func main() {
	var a int8
	_, _ = fmt.Scan(&a)
	switch a {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 4, 6, 9, 11:
		fmt.Println(30)
	default:
		fmt.Println(29)
	}
}
