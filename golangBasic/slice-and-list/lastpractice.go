package main

import "fmt"

func divide1(s string, i int) (string, string) {
	return s[:i], s[i:]
}
func reverse(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs)/2; i++ {
		temp := bs[i]
		bs[i] = bs[len(bs)-i-1]
		bs[len(bs)-i-1] = temp
	}
	return string(bs)
}
func findDiff(s []string) []string {
	new := make([]string, 0)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			new = append(new, s[i])
		}
	}
	return new
}
func bubble(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j+1] < s[j] {
				temp := s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
	return s
}
func mapFunc(f func(int) int, s []int) []int {
	for i, j := range s {
		s[i] = f(j)
	}
	return s
}
func times10(a int) int {
	return 10 * a
}
func main() {
	fmt.Println(reverse("abcd"))
	s := []string{"a", "a", "b", "b", "c", "d"}
	fmt.Println(findDiff(s))
	s1 := []int{1, 3, 5, 2, 7, 3, 4}
	fmt.Println(bubble(s1))
	fmt.Println(mapFunc(times10, s1))
}
