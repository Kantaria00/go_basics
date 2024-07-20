package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
	}
	for i, num := range arr {
		count := 0
		for j, num2 := range arr {
			if i != j && num == num2 {
				count = 1
				break
			}
		}
		if count == 0 {
			fmt.Print(num, " ")
		}
	}
}
