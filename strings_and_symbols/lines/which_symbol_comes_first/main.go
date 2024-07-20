package main

import (
	"fmt"
)

func main() {
	var str string
	res := "-1"
	_, _ = fmt.Scan(&str)
	for _, a := range str {
		if string(a) == "x" {
			res = "x"
			break
		} else if string(a) == "w" {
			res = "w"
			break
		}
	}
	fmt.Println(res)
}
