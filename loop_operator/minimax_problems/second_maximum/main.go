package main

import "fmt"

func main() {
	var n, max1, max2 int
	_, _ = fmt.Scan(&n)
	max1 = n
	for n != 0 {
		_, _ = fmt.Scan(&n)
		if max1 < n {
			max2 = max1
			max1 = n
		} else if max2 < n {
			max2 = n
		}
	}
	fmt.Println(max2)
}
