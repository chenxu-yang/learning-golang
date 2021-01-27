package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
func fibo2(n int) (posi, value int) {
	if n <= 1 {
		posi = n
		value = 1
		return
	}
	posi = n
	value = fibo(n-1) + fibo(n-2)
	return
}
func print10(n int) {
	fmt.Println(n)
	if n == 1 {

		return
	}
	print10(n - 1)
}
func facto(n int) int {

}
func main() {
	for i := 0; i < 10; i++ {
		posi, value := fibo2(i)
		fmt.Printf("fibo(%d)is%d\n", posi, value)
	}
	print10(10)
}
