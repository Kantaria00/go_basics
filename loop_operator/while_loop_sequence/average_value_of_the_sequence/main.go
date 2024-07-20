package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	sum, cnt := 0.0, 0.0
	for n != 0 {
		sum += float64(n)
		cnt++
		_, _ = fmt.Scan(&n)
	}
	fmt.Println(sum / cnt)
}
