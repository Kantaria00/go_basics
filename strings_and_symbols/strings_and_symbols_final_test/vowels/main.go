package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	var count int
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	s = sc.Text()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	fmt.Println(count)
}
