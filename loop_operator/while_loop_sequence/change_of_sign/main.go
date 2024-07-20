package main

import "fmt"

func main() {
	var n, cnt, zn int
	_, _ = fmt.Scan(&n)
	if n > 0 {
		zn = 1
	} else {
		zn = 0
	}
	for n != 0 {
		switch {
		case n > 0 && zn == 0:
			cnt++
			zn = 1
		case n < 0 && zn == 1:
			cnt++
			zn = 0
		}
		_, _ = fmt.Scan(&n)
	}
	fmt.Println(cnt)
}
