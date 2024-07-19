package main

import (
	"fmt"
)

func main() {
	var a int16
	_, _ = fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println(a + 2)
	} else {
		fmt.Println(a + 1)
	}
}
