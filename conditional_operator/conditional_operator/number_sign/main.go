package main

import (
	"fmt"
)

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	if a > 0 {
		fmt.Println(1)
	} else if a == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(-1)
	}
}
