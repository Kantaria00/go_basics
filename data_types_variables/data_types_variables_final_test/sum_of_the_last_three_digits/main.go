package main

import (
	"fmt"
)

func main() {
	var num int32
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	num %= 1000
	fmt.Println(num/100 + num%100/10 + num%10)
}
