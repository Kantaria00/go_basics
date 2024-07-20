package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n int
	ans := "YES"
	_, _ = fmt.Scan(&n)
	d := len(strconv.Itoa(n))
	for i := 1; i <= d/2; i++ {
		if n%int(math.Pow(10, float64(i)))/int(math.Pow(10, float64(i)-1.0)) != n%int(math.Pow(10, float64(d)))/int(math.Pow(10, float64(d)-1.0)) {
			ans = "NO"
			break
		}
		d--
	}
	fmt.Println(ans)
}
