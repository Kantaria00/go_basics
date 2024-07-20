package main

import "fmt"

func main() {
	var n, speed int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&speed)
	mx := speed
	mn := "NO"
	for i := 1; i < n; i++ {
		tmp := speed
		_, _ = fmt.Scan(&speed)
		if speed > mx {
			mx = speed
		}
		if speed < 30 || tmp < 30 {
			mn = "YES"
		}

	}
	fmt.Println(mx, mn)
}
