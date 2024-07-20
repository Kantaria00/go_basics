package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	if checkHappy(a) == checkHappy(b) && checkHappy(a) == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
func checkHappy(x int) int {
	if sumTicket(x/1000) == sumTicket(x%1000) {
		return 1
	} else {
		return -1
	}
}
func sumTicket(x int) int {
	var sum int
	for x != 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}
