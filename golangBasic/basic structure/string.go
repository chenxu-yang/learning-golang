package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("love", "lo"))
	fmt.Println(strings.Contains("love", "lo"))
	fmt.Println(strings.Index("love", "lo"))
	fmt.Println(strings.Count("love", "lo"))
	fmt.Println(strings.Repeat("love", 2))
	fmt.Println(strings.ToLower("LOVE"))
	fmt.Println(strings.ToUpper("love"))
	fmt.Println(strings.TrimSpace("  love  "))
	s := strings.Fields("  love can hurt  ")
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Println(strings.Join(s, " "))
	fmt.Println(strings.Split("love can hurt", " "))

	fmt.Println(strings.Replace("lovelololo", "lo", "dec", 2))

}
