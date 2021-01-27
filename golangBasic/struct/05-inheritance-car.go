package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}
type Car struct {
	Engine
	wheelCount int
}

type benci struct {
	Car
}

func (c *Car) getWheelNumber() int {
	return c.wheelCount
}
func (b *benci) say() {
	fmt.Println("hello from benci")
}
func (c *Car) Start() {
	fmt.Println("start")
}
func (c *Car) Stop() {
	fmt.Println("stop")
}
func (c *Car) work() {
	c.Start()
	c.Stop()
}
func main() {
	m := benci{Car{nil, 4}}
	fmt.Println("number of wheels:", m.getWheelNumber())
	m.work()
	m.say()

}
