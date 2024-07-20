package main

import (
	"fmt"
)

func main() {
	var c string
	var num int
	_, _ = fmt.Scan(&c, &num)
	if len(c) < num {
		fmt.Println("NO")
	} else {
		fmt.Println(string(c[num-1]))
	}
}
