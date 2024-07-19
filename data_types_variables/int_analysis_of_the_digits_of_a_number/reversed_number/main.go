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
	fmt.Print(num%10*100 + num%100/10*10 + num/100)
}
