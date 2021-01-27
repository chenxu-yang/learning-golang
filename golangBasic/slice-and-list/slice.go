package main

import (
	"fmt"
)

func main() {
	//切片本身就是一个指针
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("the length of slice1 is %d\n", len(slice1))
	fmt.Printf("the length of arr1 is %d\n", len(arr1))
	fmt.Printf("the capacity of slice1 is %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
	var list = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(list[:]))

	//make创建切片
	//s2 := make([]int, 50)
	//v:=make([]int 10 50)//分配一个有50int值得数组，并且创建了一个长度为10，容量为50的切片，指向数组前10个元素。
	s := make([]byte, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	p := s[2:4]
	fmt.Println(len(p))
	fmt.Println(cap(p))

	s1 := []byte{'p', 'a', 's', 'f'}
	s3 := s1[1:]
	fmt.Println(s3)
	s1[1] = 't'
	fmt.Println(s1)
	fmt.Println(s3)
}
func sum(a []int) int {
	sums := 0
	for _, i := range a {
		sums += i
	}
	return sums
}
