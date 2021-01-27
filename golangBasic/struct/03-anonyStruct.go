package main

import "fmt"

type a struct {
	a float64
	int
	string
}

func main() {
	struct1 := a{2.2, 8, "dsa"}
	fmt.Println(struct1)
}
