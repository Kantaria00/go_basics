package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&n)
		arr[i] = n
		if i%2 != 0 {
			tmp := arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp
			fmt.Print(arr[i-1], " ", arr[i], " ")
		}
	}
	if len(arr)%2 != 0 {
		fmt.Print(arr[len(arr)-1])
	}
}
