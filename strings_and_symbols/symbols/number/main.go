package main

import (
	"fmt"
)

func main() {
	var c rune
	_, err := fmt.Scan(&c)
	if err != nil {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
