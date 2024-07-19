package main

import (
	"fmt"
)

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	if a <= 13 {
		fmt.Println("детство")
	} else if a > 13 && a <= 24 {
		fmt.Println("молодость")
	} else if a > 24 && a <= 59 {
		fmt.Println("зрелость")
	} else {
		fmt.Println("старость")
	}
}
