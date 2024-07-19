package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var st1, st2 string
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	st1 = sc.Text()
	_ = sc.Scan()
	st2 = sc.Text()
	if st1 == st2 {
		fmt.Println("Пароль принят")
	} else {
		fmt.Println("Пароль не принят")
	}
}
