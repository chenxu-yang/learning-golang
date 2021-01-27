package main

import "fmt"

var fibos [50]int

func fibo() {
	fibos[0] = 1
	fibos[1] = 1
	for i := 2; i < 50; i++ {
		fibos[i] = fibos[i-1] + fibos[i-2]
	}
}
func main() {
	var arr1 [5]int
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i := 0; i < len(arr1); i++ {
		arr1[i] = 2 * i
	}
	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Printf("ele at index %d is %d\n", i, arr1[i])
	// }
	// for index, ele := range arr1 {
	// 	fmt.Printf("ele at index %d is %d\n", index, ele)
	// }
	// fibo()
	// fmt.Println(fibos)
	arr2 := &arr1
	fmt.Println("arr2 is ", *arr2)
	fmt.Println("arr1 is ", arr1)
	fmt.Println("location of fibos2 is", &arr2[0])
	fmt.Println("location of fibos1 is", &arr1[0])
}
