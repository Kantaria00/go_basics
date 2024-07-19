package main

import (
	"fmt"
)

func main() {
	var a int16
	_, _ = fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
