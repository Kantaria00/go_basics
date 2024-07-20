package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	cntP, cntM := 0, 0
	for n != 0 {
		switch {
		case n > 0:
			cntP++
		case n < 0:
			cntM++
		}
		_, _ = fmt.Scan(&n)
	}
	fmt.Println(cntP - cntM)
}
