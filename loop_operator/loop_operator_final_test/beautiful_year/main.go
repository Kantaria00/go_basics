package main

import (
	"fmt"
)

func main() {
	var year, flag2 int
	flag := "YES"
	_, _ = fmt.Scan(&year)
	year++
	for flag2 != 6 {
		for i := 1; i <= 1000; i *= 10 {
			for j := i * 10; j <= 1000; j *= 10 {
				if year/i%10 == year/j%10 {
					flag = "NO"
					flag2 = 0
					break
				} else {
					flag2++
				}
			}
			if flag == "NO" {
				year++
				break
			}
		}
		flag = "YES"
	}
	fmt.Println(year)
}
