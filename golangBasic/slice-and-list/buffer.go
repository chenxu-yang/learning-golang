package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*
			定义buffer:
			var byte bytes.Buffer
			var r *bytes.Buffer=new(bytes.Buffer)

		var buffer bytes.Buffer
		for {
			if s, ok := getNextString(); ok {
				buffer.WriteString(s)
			} else {
				break
			}
		}
		fmt.Print(buffer.String(), "\n")
		这种实现方式比使用+=更加节省内存的CPU
	*/
	var byte bytes.Buffer
	byte.WriteString("absfgddfg")
	fmt.Print(divide(&byte, 4))
}
func divide(buf *bytes.Buffer, n int) (a, b []byte) {
	a = buf.Bytes()[:n]
	b = buf.Bytes()[n+1:]
	return

}
