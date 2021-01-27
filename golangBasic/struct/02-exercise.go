package main

import "fmt"

type address struct {
	addr string
}
type vcard struct {
	name     string
	location *address
}
type rectangle struct {
	length int
	width  int
}

func area(l *rectangle) int {
	return l.width * l.length
}
func main() {
	addr := &address{"chongqing"}
	chenxu := &vcard{"chenxu", addr}
	fmt.Println(*chenxu.location)
	a := &rectangle{1, 2}
	fmt.Println(area(a))
}
