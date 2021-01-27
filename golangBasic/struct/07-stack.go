package main

import (
	"fmt"

	"./stack"
)

var s stack.Stack

func main() {
	s.Push(1)
	fmt.Println(s)
}
