package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	fmt.Println(reverseNum(a) + reverseNum(b))
}
func reverseNum(x int) int {
	var res int
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res
}
