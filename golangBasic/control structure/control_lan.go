package main

import (
	"fmt"
	"reflect"
	"time"
)

func loop1(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}
func loop2(n int) int {
	sum := 0
	for sum < n {
		sum += n
		if k := n * 2; k > 6 {
			fmt.Println("if")
			sum++
		} else {
			fmt.Println("else")
		}
	}
	return sum
}
func newton(n float64) int {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - n) / (2 * z)
		fmt.Println(i, z)
	}
	return 1
}
func saturday() {
	fmt.Println("when's saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("two days far")
	default:
		fmt.Println("far")
		fmt.Println(today)
		fmt.Println(reflect.TypeOf(today))
	}
}
func gettime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("a")
	case t.Hour() < 17:
		fmt.Println("b")
	default:
		fmt.Println('c')
	}
}
func usedefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
func tag() {
LABLE1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 4 {
				continue LABLE1
			}
			fmt.Printf("i is: %d,and j is: %d\n", i, j)
		}
	}
}
func gotos() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
func main() {
	// fmt.Println(loop1(10))
	// fmt.Println(loop2(10))
	// fmt.Println(hello.Add(1, 2))
	// newton(10)
	// fmt.Println("go runs on ")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("hh")
	// 	fmt.Println("os x")
	// case "linux":
	// 	fmt.Println("linux")
	// default:
	// 	fmt.Println("%s.\n", os)

	// }
	// saturday()
	// gettime()
	// usedefer()
	tag()
	gotos()
}
