package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
