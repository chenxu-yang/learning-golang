package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	bks := make([]book, 0)
	file, err := os.Open("products.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = string(line)[:len(line)-1]
		strSl := strings.Split(line, ";")
		book := new(book)
		book.title = strSl[0]
		book.price, err = strconv.ParseFloat(strSl[1], 32)
		if err != nil {
			fmt.Printf("error in file: %v", err)
		}
		book.quantity, err = strconv.Atoi(strSl[2])
		if err != nil {
			fmt.Printf("error in book: %v", err)
		}
		bks = append(bks, *book)
	}
	fmt.Println("we have read the following books from the file:")
	for _, book := range bks {
		fmt.Println(book)
	}
}
