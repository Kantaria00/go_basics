package main

import (
	"fmt"
)

func main() {
	var str, str1 string
	_, _ = fmt.Scan(&str)
	for _, a := range str {
		str1 = string(a) + str1
	}
	if str1 == str {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
