package main

import "fmt"

func bytes() {
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}

func main() {
	type length int
	var x length
	var y int
	x = 40
	y = int(x)
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	i, j := 42, 20
	p := &i
	fmt.Println(*p)
	*p = 1
	fmt.Println(i)
	q := &j
	*q = *q / 20
	fmt.Println(*q)
	bytes()

}
