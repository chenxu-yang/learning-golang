package main

import "fmt"

func main() {
	items := []int{1, 2, 3, 4}
	for _, item := range items {
		item *= 2
	}
	fmt.Println(items)
	sum := 0
	fmt.Println(add(sum, items...))
}
func add(a int, b []int) int {
	return a + b
}
