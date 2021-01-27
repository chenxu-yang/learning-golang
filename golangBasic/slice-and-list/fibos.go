package main

import "fmt"

func main() {
	fmt.Println(fiboos(10))

}
func fiboos(n int) []int {

	result := make([]int, n)
	result[0] = 1
	result[1] = 1
	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result
}
