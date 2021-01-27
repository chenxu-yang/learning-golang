package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("error")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}

	// simple data write:
	fmt.Fprintf(outputFile, “Some test data.\n”)
	outputWriter.Flush()

}
