package main

import "fmt"

func main() {
	var n, mn, mx, indMax, indMin int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
		if i == 0 || arr[i] < mn {
			mn = arr[i]
			indMin = i
		}
		if i == 0 || arr[i] > mx {
			mx = arr[i]
			indMax = i
		}
	}
	fmt.Println(indMax - indMin)
}
