package main

import "fmt"

var a string

// defer允许我们推迟到函数返回之前一刻才执行某个语句
func function1() {
	fmt.Println("top of function1")
	defer function2()
	fmt.Println("bottom of function1")
	return
}
func function2() {
	fmt.Println("function2")
}

func testAdd() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 倒序输出 类似栈
func testLoop() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
func n() { fmt.Println(a) }

func m() {
	a = "9"
	print(a)
	n()
}
func main() {
	// function1()
	// testAdd()
	// testLoop()
	// a := 1
	// b := a
	// fmt.Println(&a)
	// fmt.Println(&b)
	a = "G"
	print(a)
	m()
}
