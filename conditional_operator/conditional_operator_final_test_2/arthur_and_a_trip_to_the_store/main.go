package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Println(min(a+b+c, (a+c)*2, (b+c)*2, (a+b)*2))
}
