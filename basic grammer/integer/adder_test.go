package integer

import (
	"fmt"
	"testing"
)

func Add(x, y int) int {
	return x + y
}
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("exprected '%d' but got '%d'", expected, sum)
	}

}
