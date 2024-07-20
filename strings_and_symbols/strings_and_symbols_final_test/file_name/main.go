package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	var n, count1 int
	_, err := fmt.Scan(&n)
	if err != nil {
	}
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s = sc.Text()
	for _, i := range s {
		if i == 'x' {
			count1++
		} else {
			count1 = 0
		}
		if count1 > 2 {
			n--
		}
	}
	fmt.Println(len(s) - n)
}
