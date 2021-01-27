package main

import "fmt"

type rectangle1 struct {
	length int
	width  int
}

func (r *rectangle1) area() int {
	return r.length * r.width
}
func main() {
	a := rectangle1{2, 3}
	fmt.Println(a.area())
}
