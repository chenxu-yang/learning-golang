package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../struct/stack"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	cals := new(stack.Stack)
	fmt.Println("enter two number and an operator, or q to stop")
	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("input error")
			os.Exit(1)
		}
		token = token[:len(token)-1]
		switch {
		case token == "q":
			fmt.Println("calcu stoped")
			return

		case token >= "0" && token <= "99999":
			i, _ := strconv.Atoi(token)
			cals.Push(i)
		case token == "*":
			a, _ := cals.Pop()
			b, _ := cals.Pop()
			fmt.Printf("%d * %d equals to %d:", a, b, a*b)
		case token == "/":
			a, _ := cals.Pop()
			b, _ := cals.Pop()
			fmt.Printf("%d / %d equals to:%d:", a, b, a/b)
		case token == "+":
			a, _ := cals.Pop()
			b, _ := cals.Pop()
			fmt.Printf("%d + %d equals to:%d:", a, b, a+b)
		case token == "-":
			a, _ := cals.Pop()
			b, _ := cals.Pop()
			fmt.Printf("%d - %d equals to:%d:", a, b, a-b)
		default:
			fmt.Println("invalid input")
		}

	}

}
