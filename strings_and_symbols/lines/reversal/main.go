package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str, res string
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	str = sc.Text()
	for _, a := range str {
		res = string(a) + res
	}
	fmt.Println(res)
}
