package main

import (
	"fmt"
)

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Println("Следующее за числом", num, "число:", num+1)
	fmt.Println("Для числа", num, "предыдущее число:", num-1)
}
