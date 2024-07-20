package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			fmt.Println(i)
			break
		}
	}
}
