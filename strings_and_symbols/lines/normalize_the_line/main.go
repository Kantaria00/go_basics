package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str = sc.Text()
	for i := range str {
		if string(str[i]) == " " && string(str[i-1]) == " " && i != 0 {
			continue
		}
		fmt.Print(string(str[i]))
	}
}
