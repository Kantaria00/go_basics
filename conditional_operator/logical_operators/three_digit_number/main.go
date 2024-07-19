package main

import (
	"fmt"
)

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	if a/1000 == 0 && a/100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
