package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	chas  int
	words int
	lines int
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter something, type s to stop:")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("error!")
		}
		if input == "s\n" {
			fmt.Println("here are the counts")
			fmt.Println("number of characters:", chas)
			fmt.Println("number of words:", words)
			fmt.Println("number of lines:", lines)
			os.Exit(0)
		}
		Counters(input)
	}
}
func Counters(input string) {
	chas += len(input) - 1
	words += len(strings.Fields(input))
	lines++
}
