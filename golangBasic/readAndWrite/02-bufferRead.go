package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("enter something:")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("the input is:", input)
	}
}
