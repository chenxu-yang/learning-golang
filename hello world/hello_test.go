// testing function, rules:
// write in the file named:xxx_test.go
// name of function start with 'Test'
// only one input: t *testing.T t是测试框架的hook
package main

import "testing"

// func TestHello(t *testing.T){
//	got :=hello("chenxu")
//	want:="hello,chenxu"
//	if got !=want{
//		t.Errorf("got '%q' want ''%q",got,want)
//	}
//}

// 子测试 对一个事情进行分组测试 测试不同情况
func TestHello(t *testing.T) {
	//t.Run("saying hello to people", func(t *testing.T) {
	//	got :=hello("chenxu")
	//	want:="hello,chenxu"
	//	if got !=want {
	//		t.Errorf("got '%q' want '%q",got,want)
	//	}
	//})
	//t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
	//	got:=hello("")
	//	want:="hello,world"
	//	if got!=want{
	//		t.Errorf("got '%q' want '%q',got,want")
	//	}
	//
	//})
	// 重构
	assertmessage := func(t *testing.T, got, want string) {
		t.Helper() //辅助函数，测试失败时报告的行号在函数调用中而不是在辅助函数内部
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("chenxu", "")
		want := "hello,chenxu"
		assertmessage(t, got, want)
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := hello("", "")
		want := "hello,world"
		assertmessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := hello("elodie", "spanish")
		want := "hola,elodie"
		assertmessage(t, got, want)
	})
}

//notes:
