package main

import "fmt"

func main() {
	var a, b, c int
	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Println(Fact2(a), Fact2(b), Fact2(c))
}
func Fact2(n int) int {
	if n%2 == 0 {
		return factorial(n, 2)
	} else {
		return factorial(n, 1)
	}
}
func factorial(n int, k int) int {
	var res = k
	for i := k + 2; i <= n; i += 2 {
		res *= i
	}
	return res
}
