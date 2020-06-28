package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

/*源码：
It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
type Writer interface {
   Write(p []byte) (n int, err error)
}*/

//io.Writer 是一个很好的通用接口，用于将数据放在某个地方
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} // bytes包中的buffer类型实现了writer接口
	Greet(&buffer, "Chris")
	got := buffer.String()
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
