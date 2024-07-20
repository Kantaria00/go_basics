package main

import "fmt"

func main() {
	var n, m int
	_, _ = fmt.Scan(&n)
	mxA := findMaximum(n)
	_, _ = fmt.Scan(&m)
	mxB := findMaximum(m)
	fmt.Println(mxA * mxB)
}
func findMaximum(num int) int {
	var mx, a int
	_, _ = fmt.Scan(&a)
	mx = a
	for i := 1; i < num; i++ {
		_, _ = fmt.Scan(&a)
		if a > mx {
			mx = a
		}
	}
	return mx
}
