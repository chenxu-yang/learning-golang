package struct__method_and_interface

import (
	"testing"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Width  float64
	Height float64
}

//声明方法 针对不同结构
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (t Triangle) Area() float64 {
	return 0.5 * t.Height * t.Width
}

// 测试计算周长函数
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if want != got {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//测试计算面积函数
//func TestArea(t *testing.T){
//checkArea:=func(t *testing.T,shape Shape,want float64){
//	t.Helper()
//	got:=shape.Area()
//	if want!=got{
//		t.Errorf("got %.2f want %.2f",got,want)
//	}
//}
//t.Run("rectangles", func(t *testing.T) {
//	rectangle:=Rectangle{12.0,6.0}
//	checkArea(t,rectangle,72.0)
//})
//t.Run("circles", func(t *testing.T) {
//	circle:=Circle{10}
//	checkArea(t,circle,314.0)
//})
//}
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{"Circle", Circle{10}, 314},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got %.2f want %.2f", got, tt.hasArea)
		}
	}
}

// method is like functions,
// func (receiverName ReceiverType)MethodName(args)
// 把类型的第一个字母作为接受者是go的惯例

//interface
// type Shape interface

//总结
//声明结构创建自己类型，
//声明接口，定义适合不同参数类型的函数，比如面积函数针对不同形状
//创建不同的方法来实现接口
