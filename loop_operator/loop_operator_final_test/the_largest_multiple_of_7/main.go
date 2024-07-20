package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	mx := -20000
	flag := "NO"
	for i := a; i <= b; i++ {
		if i > mx && i%7 == 0 {
			mx = i
			flag = "YES"
		}
	}
	switch {
	case flag == "NO":
		fmt.Println(flag)
	default:
		fmt.Println(mx)
	}
}
