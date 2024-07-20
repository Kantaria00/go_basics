package main

import "fmt"

func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
	}
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '.':
			fmt.Print(0)
		case s[i] == '-' && s[i+1] == '.':
			fmt.Print(1)
			i++
		case s[i] == '-' && s[i+1] == '-':
			fmt.Print(2)
			i++
		}
	}
}
