package main

import (
	"errors"
	"fmt"
	"math"
)

func mult1(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}
func mult2(a, b int) (x1, x2, x3 int) {
	x1 = a + b
	x2 = a * b
	x3 = a - b
	return
}
func sqrt1(x float64) (float64, error) {
	if x < 0 {
		return float64(math.NaN()), errors.New("nagative")
	}
	return math.Sqrt(x), nil
}
func mult3(a, b int, n *int) {
	*n = a * b
}
func longs(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
func longPrint(s ...int) {
	for _, v := range s {
		fmt.Println("value is %v\n", v)
	}
}
func main() {
	fmt.Println(mult1(1, 2))
	fmt.Println(mult2(1, 2))
	fmt.Println(sqrt1(-1))
	b := 0
	n := &b
	mult3(1, 2, n)
	fmt.Println(*n)
	fmt.Println(longs(2, 3, 4, -1))
	longPrint(2, 3, 4, -1, 9, 0, 6)
}
