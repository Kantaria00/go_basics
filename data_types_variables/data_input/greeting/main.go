package main

import (
	"fmt"
)

func main() {
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	fmt.Println("Привет,", name)
}
