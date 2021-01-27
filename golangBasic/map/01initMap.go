package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	value, isPresent := map1["c"]
	fmt.Println(value, isPresent)
	fmt.Println(map1)
	delete(map1, "a")
	fmt.Println(map1)
}
