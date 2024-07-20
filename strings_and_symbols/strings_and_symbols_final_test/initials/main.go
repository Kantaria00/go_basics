package main

import (
	"fmt"
)

func main() {
	var st1, st2, st3 string
	_, _ = fmt.Scan(&st1, &st2, &st3)
	fmt.Print(st1, " ", string([]rune(st2)[0]), ".", string([]rune(st3)[0]), ".")
}
