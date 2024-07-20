package main

import (
	"fmt"
)

func main() {
	var n, count, c3, c2, c1 int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
		switch {
		case arr[i] == 4:
			count++
		case arr[i] == 3:
			if c1 != 0 {
				count++
				c1--
			} else {
				c3++
			}
		case arr[i] == 2:
			if c2 != 0 {
				count++
				c2 = 0
			} else {
				c2++
			}
		case arr[i] == 1:
			if c3 != 0 {
				count++
				c3--
			} else {
				c1++
			}
		}
	}
	switch {
	case c1 == 0:
		fmt.Println(count + c2 + c3)
	case c2 == 0:
		if c1%4 == 0 {
			fmt.Println(count + c1/4 + c3)
		} else {
			fmt.Println(count + c1/4 + 1 + c3)
		}
	default:
		if (c1+c2)%4 == 0 {
			fmt.Println(count + (c1+c2)/4 + c3)
		} else {
			fmt.Println(count + (c1+c2)/4 + 1 + c3)
		}
	}
}
