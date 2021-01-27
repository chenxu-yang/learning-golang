package main

import "fmt"

type Shaper interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side float64
}
type Rectangle struct {
	length float64
	width  float64
}
type Triangle struct {
	width float64
	hight float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}
func (s *Square) Area() float64 {
	return s.side * s.side
}
func (t *Triangle) Area() float64 {
	return t.width * t.hight / 2
}
func getArea(shape Shaper) float64 {
	return shape.Area()
}
func main() {
	s := &Square{2}
	fmt.Println("the area of this square is:", getArea(s))
	r := &Rectangle{2, 3}
	fmt.Println("the area of this rectangle is:", getArea(r))
	t := &Triangle{2, 5}
	fmt.Println("the area of this triangle is:", getArea(t))

}
