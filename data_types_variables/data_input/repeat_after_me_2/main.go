package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	a := scanner.Text()
	_ = scanner.Scan()
	b := scanner.Text()
	_ = scanner.Scan()
	c := scanner.Text()

	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(a)
}
