package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		name := scanner.Text()
		fmt.Println(name)
	}
}
