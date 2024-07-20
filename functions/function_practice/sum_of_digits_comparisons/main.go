package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	res := sumNum(a) - sumNum(b)
	switch {
	case res > 0:
		fmt.Println(1)
	case res < 0:
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}
func sumNum(num int) int {
	var sum int
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
