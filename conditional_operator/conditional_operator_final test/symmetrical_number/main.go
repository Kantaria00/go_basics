package main

import (
	"fmt"
)

func main() {
	var a int16
	_, _ = fmt.Scan(&a)

	if a/1000 == a%10 && a%1000/100 == a%100/10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
