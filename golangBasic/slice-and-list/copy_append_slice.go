package main

import "fmt"

func broad(factor int, slice []int) []int {
	new_slice := make([]int, factor*len(slice))
	copy(new_slice, slice)
	return new_slice
}
func insertStringSlice(s1, s2 []int, position int) []int {
	// n := len(s1) + len(s2)
	// newSlice := make([]int, position, n)
	// copy(newSlice, s1[:position])
	// newSlice = append(newSlice, s2...)
	// newSlice = append(newSlice, s1[position:]...)
	// return newSlice
	return append(append(s1[:position], s2...), s1[position:]...)
}
func removeStringSlice(s1 []int, start, end int) []int {
	// newSlice := make([]int, start)
	// copy(newSlice, s1[:start])
	// return append(newSlice, s1[end+1:]...)
	return append(s1[:start], s1[end+1:]...)

}
func main() {
	// s1 := []int{1, 2, 3}
	// s2 := make([]int, 10)
	// n := copy(s2, s1)
	// fmt.Println(s1, s2)
	// fmt.Println(len(s2), cap(s2))
	// fmt.Println(n)
	// s3 := []int{4, 5, 6}
	// s3 = append(s3, 7, 8, 9)
	// fmt.Println(s3)
	// fmt.Println(broad(2, s1))
	s4 := []int{1, 2, 3}
	s5 := []int{4, 5, 6}
	s6 := make([]int, 3, 10)
	copy(s6, s5)
	fmt.Println("s6:,", s6)
	fmt.Println(insertStringSlice(s4, s5, 2))
	fmt.Println(insertStringSlice(s6, s4, 1))
	s7 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(removeStringSlice(s7, 2, 3))
}
