package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	_, err := fmt.Scan(&num1, &num2, &num3)
	if err != nil {
		return
	}
	fmt.Println(num1+num2+num3, num1*num2*num3)
}
