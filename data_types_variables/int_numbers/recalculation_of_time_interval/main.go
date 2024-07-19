package main

import (
	"fmt"
)

func main() {
	var time int
	_, err := fmt.Scan(&time)
	if err != nil {
		return
	}
	fmt.Println(time, "мин - это", time/60, "час", time%60, "минут.")
}
