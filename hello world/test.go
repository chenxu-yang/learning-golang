//package combine different go code together
//:=是声明变量并赋值
package main

import "fmt"

//func main()  {
//	fmt.Println("hello,world")
//}
const englishhelloprefix = "hello,"
const spanishhelloprefix = "hola,"
const frenchhelloprefix = "bonjour,"

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchhelloprefix
	case "spanish":
		prefix = spanishhelloprefix
	default:
		prefix = englishhelloprefix
	}
	return
}
func main() {
	fmt.Println(hello("chenxu", ""))
}
