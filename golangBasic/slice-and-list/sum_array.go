package main

import "fmt"

func main() {
	s := []float64{1.2, 1.3, 1.4, 1.5}
	fmt.Println(Sum(s))
	s1 := []int{1, 5, 32, 2, -1}
	fmt.Println(minmax(s1))
	s2 := s1[0:3]
	fmt.Println(len(s2))
	fmt.Println(len(s[1:2]))
}
func Sum(arr []float64) (int, float64) {
	sum := 0.0
	for _, item := range arr {
		sum += item
	}
	return int(sum), sum
}
func minmax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	curMin, curMax := arr[0], arr[0]
	for _, item := range arr {
		if item < curMin {
			curMin = item
		}
		if item > curMax {
			curMax = item
		}
	}
	return curMin, curMax
}
