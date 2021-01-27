package main

import "fmt"

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}
type Cars []*Car

func main() {
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("all cars:", allCars)
	fmt.Println("new bmws:", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	fmt.Println("Map sortedCars:", sortedCars)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars:", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("we have", BMWCount, "BMWS")
}

func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(car *Car) {
		if f(car) {
			cars = append(cars, car)
		}
	})
	return cars
}
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	sortedCars["test"] = make([]*Car, 0)
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	appender := func(car *Car) {
		if _, ok := sortedCars[car.Manufacturer]; ok {
			sortedCars[car.Manufacturer] = append(sortedCars[car.Manufacturer], car)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], car)

		}
	}
	return appender, sortedCars
}
