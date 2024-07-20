package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	var res int
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	str = sc.Text()
	for _, a := range str {
		if string(a) == " " {
			res++
		}
	}
	fmt.Println(res + 1)
}
