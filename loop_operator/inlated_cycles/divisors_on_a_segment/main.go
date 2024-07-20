package main

import "fmt"

func main() {
	var a, b, k int
	_, _ = fmt.Scan(&a, &b, &k)
	for i := a; i <= b; i++ {
		temp := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				temp++
			}
		}
		if temp >= k {
			fmt.Print(i, " ")
		}
	}
}
