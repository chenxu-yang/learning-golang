package main

import "fmt"

type Shaper interface {
	Area() float64
}
type Square struct {
	side float64
}
type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}
func (s *Square) Area() float64 {
	return s.side * s.side
}
func main() {
	s := &Square{2}
	r := &Rectangle{2, 3}
	shapes := []Shaper{s, r}
	for n := range shapes {
		fmt.Println("shape:", shapes[n])
		fmt.Println("area is:", shapes[n].Area())
	}
}
