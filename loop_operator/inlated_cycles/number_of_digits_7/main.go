package main

import "fmt"

func main() {
	var N, count int
	_, _ = fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		temp := i
		for temp != 0 {
			if temp%10 == 7 {
				count++
			}
			temp /= 10
		}
	}
	fmt.Println(count)
}
