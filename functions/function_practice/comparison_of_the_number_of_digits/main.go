package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	res := countNum(a) - countNum(b)
	switch {
	case res > 0:
		fmt.Println(1)
	case res < 0:
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}
func countNum(num int) int {
	var count int
	for num != 0 {
		count++
		num /= 10
	}
	return count
}
