package main

import (
	"fmt"
)

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Println(num)
	num++
	fmt.Println(num)
	num++
	fmt.Println(num)
}
