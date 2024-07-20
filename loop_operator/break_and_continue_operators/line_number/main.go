package main

import "fmt"

func main() {
	var s, ch string
	nb := 1
	_, _ = fmt.Scan(&s)
	for {
		_, _ = fmt.Scan(&ch)
		if ch == s {
			fmt.Println(nb)
			break
		}
		nb++
	}
}
