package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var st1, st2 string
	reader := bufio.NewReader(os.Stdin)
	st1, _ = reader.ReadString('\n')
	st2, _ = reader.ReadString('\n')
	if []rune(st1)[0] == []rune(st2)[len(st2)-2] && st1 != "\n" && st2 != "\n" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
