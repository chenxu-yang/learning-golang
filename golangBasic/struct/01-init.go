package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float64
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 2
	ms.f1 = 2.2
	ms.str = "Chris"
	fmt.Println(ms)
	ms2 := &struct1{2, 33.3, "test"}
	fmt.Println(ms2)
	ms3 := struct1{2, 33.3, "test"}
	fmt.Println(ms3)
}

/*
type myStruct struct { i int }
var v myStruct    // v是结构体类型变量
var p *myStruct   // p是指向一个结构体类型变量的指针
v.i
p.i
*/
