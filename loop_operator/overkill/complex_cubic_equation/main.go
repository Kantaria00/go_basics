package main

import "fmt"

func main() {
	var a, b, c, d, e, count int
	_, _ = fmt.Scan(&a, &b, &c, &d, &e)
	for x := 0; x <= 1000; x++ {
		if float64(a*x*x*x+b*x*x+c*x+d)/float64(x-e) == 0.0 {
			count++
		}
	}
	fmt.Println(count)
}
