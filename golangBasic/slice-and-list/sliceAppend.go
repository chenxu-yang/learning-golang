package main

import "fmt"

func main() {
	s := []byte{'a', 'b', 'c'}
	data := []byte{'s', 'd', 'e', 'f'}
	fmt.Printf("cap of s is %d\n", cap(s))
	fmt.Printf("len of s is %d\n", len(s))
	fmt.Println(appends(s, data))

}

/*
    ----------------
	 ...的用法：
	 1.用于函数有多个不定参数的情况
	 func test1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
	}
	2.slice可以被打散进行传递
	----------------
*/

func appends(s, data []byte) []byte {
	if len(s)+len(data) <= cap(s) {
		s = append(s, data...)
	} else {
		newS := make([]byte, len(s), len(data)+len(s))
		copy(newS, s)
		fmt.Println("now new S is:", newS)
		fmt.Printf("cap of ns is %d\n", cap(newS))
		fmt.Printf("len of ns is %d\n", len(newS))
		s = append(newS, data...)
	}
	return s

}
