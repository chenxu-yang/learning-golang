package stack

import (
	"strconv"
)

type Celsius float64

func (c Celsius) String() string {
	return "The temperature is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"

}

type Stack [4]int

func (s *Stack) Push(a int) {
	for i, v := range s {
		if v == 0 {
			s[i] = a
			break
		}
	}
}
func (s *Stack) Pop() (int, error) {
	var err error
	value := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 0 {
			value = s[i]
			s[i] = 0
			break
		}
	}
	return value, err
}
func (s *Stack) String() string {
	str := ""
	for i, v := range s {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}

type stack1 struct {
	index int
	data  [4]int
}

func (s *stack1) push(a int) {
	if s.index > 4 {
		return
	}
	s.data[s.index] = a
	s.index++
}
func (s *stack1) pop() {
	s.data[s.index-1] = 0
	s.index--
}
func (s *stack1) String() string {
	str := ""
	for i, v := range s.data {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}

// func main() {
// 	// var c Celsius = 18.36
// 	// fmt.Println(c)
// 	// stack1 := new(stack)
// 	// fmt.Println(stack1)
// 	// stack1.push(5)
// 	// fmt.Println(stack1)
// 	// stack1.push(4)
// 	// fmt.Println(stack1)
// 	// stack1.push(5)
// 	// fmt.Println(stack1)
// 	// stack1.push(5)
// 	// fmt.Println(stack1)
// 	// stack1.pop()
// 	// fmt.Println(stack1)
// 	stack1 := new(stack1)
// 	fmt.Println(stack1)
// 	stack1.push(5)
// 	fmt.Println(stack1)
// 	stack1.push(4)
// 	fmt.Println(stack1)
// 	stack1.push(5)
// 	fmt.Println(stack1)
// 	stack1.push(5)
// 	fmt.Println(stack1)
// 	stack1.pop()
// 	fmt.Println(stack1)
// }
