package main

import "fmt"

func main() {
	var a, b, c int
	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Println(sumRange(a, b) + sumRange(b, c))
}
func sumRange(x, y int) int {
	var sum int
	if x > y {
		return 0
	}
	for i := x; i <= y; i++ {
		sum += i
	}
	return sum
}
