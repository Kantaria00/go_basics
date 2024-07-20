package main

import (
	"fmt"
)

func main() {
	var str string
	_, _ = fmt.Scan(&str)
	for i := range str {
		for j := range str {
			if str[i] == str[j] && i != j {
				fmt.Println(string(str[i]))
				return
			}
		}
	}
}
