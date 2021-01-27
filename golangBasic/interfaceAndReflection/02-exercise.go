package main

import "fmt"

type simpler interface {
	Get() int
	Set(a int)
}
type rsimple struct {
	data int
}
type simple struct {
	data int
}

func (s *simple) Get() int  { return s.data }
func (s *simple) Set(a int) { s.data = a }

func (s *rsimple) Get() int  { return s.data }
func (s *rsimple) Set(a int) { s.data = a }

func test(s simpler) int {
	a := s.Get()
	fmt.Println("current:", a)
	s.Set(100)
	return s.Get()

}

func fi(s simpler) string {
	switch s.(type) {
	case *simple:
		return "simple"
	case *rsimple:
		return "rsimple"
	}
	return "unknown"
}
func main() {
	var s1 simpler
	s := &(simple{3})
	s1 = s
	fmt.Println(test(s1))
	fmt.Println(fi(s))
}
