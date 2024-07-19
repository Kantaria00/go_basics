package main

import (
	"fmt"
)

func main() {
	var num int32
	_, _ = fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
