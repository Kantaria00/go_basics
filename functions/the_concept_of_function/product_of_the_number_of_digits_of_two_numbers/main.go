package main

import "fmt"

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	fmt.Println(numberOfDigits(n) * numberOfDigits(m))
}
func numberOfDigits(x int) int {
	var num int
	for x != 0 {
		num++
		x /= 10
	}
	return num
}
