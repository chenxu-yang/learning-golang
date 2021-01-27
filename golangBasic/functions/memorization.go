package main

import (
	"fmt"
	"time"
)

func fibo1(n int) int {
	if n <= 1 {
		return 1
	}
	return fibo1(n-1) + fibo1(n-2)
}

var record [25]uint64

func fibo2(n int) uint64 {
	var result uint64
	if record[n] != 0 {
		return record[n]
	}
	if n <= 1 {
		result = 1

	} else {
		result = fibo2(n-1) + fibo2(n-2)
	}
	record[n] = result
	return result
}
func main() {
	start := time.Now()
	for i := 0; i < 25; i++ {
		fmt.Printf("the %dth fibo is %d\n", i, fibo1(i))
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("total cost time of fibo1 is %s\n", delta)

	start1 := time.Now()
	for i := 0; i < 25; i++ {
		fmt.Printf("the %dth fibo is %d\n", i, fibo2(i))
	}
	end1 := time.Now()
	delta1 := end1.Sub(start1)
	fmt.Printf("total cost time of fibo2 is %s\n", delta1)
}
