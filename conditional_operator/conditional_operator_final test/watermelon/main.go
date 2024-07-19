package main

import "fmt"

func main() {
	var w int
	_, err := fmt.Scan(&w)
	if err != nil {
		return
	}

	switch {
	case w%2 == 0 && w != 2:
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}
