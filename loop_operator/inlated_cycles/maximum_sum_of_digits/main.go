package main

import "fmt"

func main() {
	var n, sum, tempSum, resNum int
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		temp := i
		for temp != 0 {
			tempSum += temp % 10
			temp /= 10
		}
		if tempSum > sum {
			resNum = i
			sum = tempSum
		}
		tempSum = 0
	}
	fmt.Println(resNum, sum)
}
