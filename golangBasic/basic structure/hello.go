package main

import (
	"fmt"
	"math"
)

const (
	pi    = math.Pi
	big   = 1 << 100
	small = big >> 99
)

// Add two number
func Add(x, y int) int {
	return x + y
}
func swap(s, k string) (string, string) {
	return k, s
}
func split(a, b int) (x, y int) {
	x = a - 1
	y = b + 1
	return
}
func value(a int, b float64, c string) (int, int, string) {
	var d int
	d = a + 1
	s := "d"
	return d, int(b * pi), s
}
func function() {
	fmt.Println(math.Pi)
	fmt.Println(Add(1, 2))
	var a, b = swap("dsf", "a")
	fmt.Println(a, b)
	a, b = swap("dsfsafs", "a")
	// a, b = split(2, 2)
	q, w, e := value(2, 2.0, "s")
	fmt.Println(q, w, e)
}
