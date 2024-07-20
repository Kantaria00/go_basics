package main

import "fmt"

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	fN := factorial(n)
	fK := factorial(k)
	fNK := factorial(n - k)
	fmt.Println(fN / (fK * fNK))
}
func factorial(number int) int {
	fact := 1
	for i := 1; i <= number; i++ {
		fact = fact * i
	}
	return fact
}
