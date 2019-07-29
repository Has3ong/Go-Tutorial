package main

import (
	"fmt"
)

type Car struct {
	name  string
	speed int
}

func Check(i interface{}) {
	switch i.(type) {
	case Car:
		object := i.(Car)
		fmt.Println("Car ", object.name, object.speed)

	case *Car:
		object := i.(*Car)
		fmt.Println("*Car ", object.name, object.speed)

	default:
		fmt.Println("Error")
	}
}

func main() {

	Check(Car{name: "Benz", speed: 230})
	Check(&Car{name: "BMW", speed: 240})

	var Ford = new(Car)
	Ford.name = "Ford"
	Ford.speed = 200

	Check(Ford)
}
