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
		case 'e':

			fmt.Print(string('i'))
		default:

			fmt.Print(string(s[i]))
		}
	}
}
