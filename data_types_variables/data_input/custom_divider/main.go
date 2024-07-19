package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	a := sc.Text()
	_ = sc.Scan()
	b := sc.Text()
	_ = sc.Scan()
	c := sc.Text()
	_ = sc.Scan()
	d := sc.Text()
	fmt.Print(b, a, c, a, d)
}
