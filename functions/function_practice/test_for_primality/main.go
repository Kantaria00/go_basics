package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println(primality(n))
}
func primality(n int) string {
	res := "prime"
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return "composite"
		}
	}
	return res
}
