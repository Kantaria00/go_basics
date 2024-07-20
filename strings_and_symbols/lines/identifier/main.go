package main

import (
	"fmt"
)

func main() {
	var str string
	res := "YES"
	_, _ = fmt.Scan(&str)
	for i := range str {
		if !(str[0] >= 65 && str[0] <= 90 || str[0] == 95 || str[0] >= 97 && str[0] <= 122) {
			res = "NO"
			break
		}
		if !(str[i] >= 65 && str[i] <= 90 || str[i] == 95 || str[i] >= 97 && str[i] <= 122 || str[i] >= 48 && str[i] <= 57) {
			res = "NO"
			break
		}
	}
	fmt.Println(res)
}
