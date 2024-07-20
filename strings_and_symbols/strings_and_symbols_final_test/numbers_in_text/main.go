package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	s = sc.Text()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fmt.Print(string(s[i]), " ")
		}
	}
}
