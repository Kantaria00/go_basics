package main

import "fmt"

func main() {
	var n int
	flag := "NO"
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	_, _ = fmt.Scan(&n)
	arr[0] = n
	for i := 1; i < len(arr); i++ {
		_, _ = fmt.Scan(&n)
		arr[i] = n
		if arr[i-1] > 0 && arr[i] > 0 || arr[i-1] < 0 && arr[i] < 0 {
			flag = "YES"
			break
		}
	}
	fmt.Println(flag)
}
