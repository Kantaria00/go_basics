package main

import "fmt"

func main() {
	var s string
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
	}
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&s)
		if err != nil {
		}
		if len(s) > 10 {
			fmt.Print(string(s[0]), len(s)-2, string(s[len(s)-1]), "\n")
		} else {
			fmt.Println(s)
		}
	}
}
