package main

import (
	"fmt"
)

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	switch {
	case x <= 3:
		fmt.Println("начинающий")
	case x >= 4 && x <= 7:
		fmt.Println("младший разработчик")
	case x >= 8 && x <= 15:
		fmt.Println("средний разработчик")
	case x >= 16:
		fmt.Println("старший разработчик")
	}
}
