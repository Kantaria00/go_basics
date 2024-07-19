package main

import (
	"fmt"
)

func main() {
	var k, m int16
	_, _ = fmt.Scan(&k, &m)
	if k%m == 0 {
		fmt.Println(k / m)
	} else {
		fmt.Println(k/m + 1)
	}
}
