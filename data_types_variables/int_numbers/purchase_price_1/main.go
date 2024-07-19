package main

import (
	"fmt"
)

func main() {
	var num, res int
	for i := 0; i < 4; i++ {
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		res += num
	}
	fmt.Println(res * 3)
}
