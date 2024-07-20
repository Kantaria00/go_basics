package main

import "fmt"

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	fmt.Println(averageNum(n) + averageNum(m))
}
func averageNum(x int) float64 {
	var sum int
	for i := 1; i <= x; i++ {
		sum += i
	}
	return float64(sum) / float64(x)
}
