package main

import (
	"fmt"
	"time"
)

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
func fibo() func() int {
	pre := 0
	cur := 1
	return func() int {
		temp := pre + cur
		pre = cur
		cur = temp
		return temp
	}
}
func main() {
	// sum := func(n int) int {
	// 	value := 0
	// 	for i := 0; i < n; i++ {
	// 		value += i
	// 	}
	// 	return value
	// }
	// fmt.Printf("sum of %d is %d\n", 5, sum(5))
	// fv := func() {
	// 	fmt.Println("hello world")
	// }
	// fv()
	// fmt.Printf("function's value is %v", fv)
	start := time.Now()
	fmt.Println(f())
	test1 := add2()
	test2 := adder(2)
	fmt.Printf("function's value is %d\n", test1(3))
	fmt.Printf("2+3=%d\n", test2(3))
	fibos := fibo()
	for i := 0; i < 10; i++ {
		fmt.Printf("the %dth number is %d\n", i, fibos())
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}
