package main

import "fmt"

type Car struct {
	name  string
	speed int
}

type Behavior interface {
	Drive()
	Parking()
}

func (c Car) Drive() {
	fmt.Println(c.speed, " Driving")
}

func (c Car) Parking() {
	fmt.Println(c.name, " Parking")
}

func Print(i interface{}) {
	fmt.Println(i)
}

func main() {
	Benz := Car{"Benz", 230}
	Benz.Drive()
	Benz.Parking()

	Print(Benz)
}
